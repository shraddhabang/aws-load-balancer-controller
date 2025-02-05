package elbv2

import (
	"context"
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/config"
	elbv2equality "sigs.k8s.io/aws-load-balancer-controller/pkg/equality/elbv2"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/model/core"
	elbv2model "sigs.k8s.io/aws-load-balancer-controller/pkg/model/elbv2"
	"strconv"
)

// NewListenerRuleSynthesizer constructs new listenerRuleSynthesizer.
func NewListenerRuleSynthesizer(elbv2Client services.ELBV2, taggingManager TaggingManager,
	lrManager ListenerRuleManager, logger logr.Logger, featureGates config.FeatureGates, stack core.Stack) *listenerRuleSynthesizer {
	return &listenerRuleSynthesizer{
		elbv2Client:    elbv2Client,
		lrManager:      lrManager,
		logger:         logger,
		featureGates:   featureGates,
		taggingManager: taggingManager,
		stack:          stack,
	}
}

type listenerRuleSynthesizer struct {
	elbv2Client    services.ELBV2
	lrManager      ListenerRuleManager
	featureGates   config.FeatureGates
	logger         logr.Logger
	taggingManager TaggingManager

	stack core.Stack
}

func (s *listenerRuleSynthesizer) Synthesize(ctx context.Context) error {
	var resLRs []*elbv2model.ListenerRule
	s.stack.ListResources(&resLRs)
	resLRsByLSARN, err := mapResListenerRuleByListenerARN(resLRs)
	if err != nil {
		return err
	}

	var resLSs []*elbv2model.Listener
	s.stack.ListResources(&resLSs)
	for _, resLS := range resLSs {
		lsARN, err := resLS.ListenerARN().Resolve(ctx)
		if err != nil {
			return err
		}
		resLRs := resLRsByLSARN[lsARN]
		if err := s.synthesizeListenerRulesOnListener(ctx, lsARN, resLRs); err != nil {
			return err
		}
	}
	return nil
}

func (s *listenerRuleSynthesizer) PostSynthesize(ctx context.Context) error {
	// nothing to do here.
	return nil
}

func (s *listenerRuleSynthesizer) synthesizeListenerRulesOnListener(ctx context.Context, lsARN string, resLRs []*elbv2model.ListenerRule) error {
	sdkLRs, err := s.findSDKListenersRulesOnLS(ctx, lsARN)
	var lastAvailablePriority int32 = 50000
	if err != nil {
		return err
	}
	// Find rules which are matching priority
	matchedResAndSDKLRsByPriority, unmatchedResLRs, unmatchedSDKLRs := matchResAndSDKListenerRules(resLRs, sdkLRs)
	// Push down all the unmatched existing rules on load balancer
	if len(unmatchedSDKLRs) != 0 {
		p, err := s.lrManager.SetRulePriorities(ctx, unmatchedSDKLRs, lastAvailablePriority)
		if err != nil {
			return err
		}
		lastAvailablePriority = p
	}
	// Create all the unmatched/ new rules on the LB
	for _, resLR := range unmatchedResLRs {
		lrStatus, err := s.lrManager.Create(ctx, resLR, nil)
		if err != nil {
			return err
		}
		resLR.SetStatus(lrStatus)
	}
	// Filter the rules on the same priority based on their setting (actions and conditions)
	matchedResAndSDKLRsBySettings, unmatchedResAndSDKLRsBySettings, unmatchedResLRsDesiredActionsAndConditions, err := s.matchResAndSDKListenerRulesBySettings(matchedResAndSDKLRsByPriority)
	if err != nil {
		return err
	}

	var unmatchedSDKLRsBySettings []ListenerRuleWithTags
	var unmatchedResLRsBySettings []*elbv2model.ListenerRule
	if len(unmatchedResAndSDKLRsBySettings) != 0 {
		for _, resAndSDKLR := range unmatchedResAndSDKLRsBySettings {
			unmatchedSDKLRsBySettings = append(unmatchedSDKLRsBySettings, resAndSDKLR.sdkLR)
			unmatchedResLRsBySettings = append(unmatchedResLRsBySettings, resAndSDKLR.resLR)
		}
		// Push down all the unmatched existing rules by settings on load balancer
		if _, err := s.lrManager.SetRulePriorities(ctx, unmatchedSDKLRsBySettings, lastAvailablePriority); err != nil {
			return err
		}
		// Create new rules on the same priorities with new settings
		for i, resLR := range unmatchedResLRsBySettings {
			lrStatus, err := s.lrManager.Create(ctx, resLR, unmatchedResLRsDesiredActionsAndConditions[i])
			if err != nil {
				return err
			}
			resLR.SetStatus(lrStatus)
		}
	}

	// Update tags for all the matching rules
	for _, resAndSDKLR := range matchedResAndSDKLRsBySettings {
		lsStatus, err := s.lrManager.Update(ctx, resAndSDKLR.resLR, resAndSDKLR.sdkLR)
		if err != nil {
			return err
		}
		resAndSDKLR.resLR.SetStatus(lsStatus)
	}

	// Once all the required rules are created and updated, delete all the rules which were pushed down since they are not required anymore
	for _, sdkLR := range unmatchedSDKLRs {
		if err := s.lrManager.Delete(ctx, sdkLR); err != nil {
			return err
		}
	}
	for _, sdkLR := range unmatchedSDKLRsBySettings {
		if err := s.lrManager.Delete(ctx, sdkLR); err != nil {
			return err
		}
	}
	return nil
}

// findSDKListenersRulesOnLS returns the listenerRules configured on Listener.
func (s *listenerRuleSynthesizer) findSDKListenersRulesOnLS(ctx context.Context, lsARN string) ([]ListenerRuleWithTags, error) {
	sdkLRs, err := s.taggingManager.ListListenerRules(ctx, lsARN)
	if err != nil {
		return nil, err
	}
	nonDefaultRules := make([]ListenerRuleWithTags, 0, len(sdkLRs))
	for _, rule := range sdkLRs {
		if awssdk.ToBool(rule.ListenerRule.IsDefault) {
			continue
		}
		nonDefaultRules = append(nonDefaultRules, rule)
	}
	return nonDefaultRules, nil
}

type resAndSDKListenerRulePair struct {
	resLR *elbv2model.ListenerRule
	sdkLR ListenerRuleWithTags
}

type resLRDesiredActionsAndConditionsPair struct {
	desiredActions    []types.Action
	desiredConditions []types.RuleCondition
}

func matchResAndSDKListenerRules(resLRs []*elbv2model.ListenerRule, sdkLRs []ListenerRuleWithTags) ([]resAndSDKListenerRulePair, []*elbv2model.ListenerRule, []ListenerRuleWithTags) {
	var matchedResAndSDKLRs []resAndSDKListenerRulePair
	var unmatchedResLRs []*elbv2model.ListenerRule
	var unmatchedSDKLRs []ListenerRuleWithTags

	resLRByPriority := mapResListenerRuleByPriority(resLRs)
	sdkLRByPriority := mapSDKListenerRuleByPriority(sdkLRs)
	resLRPriorities := sets.Int32KeySet(resLRByPriority)
	sdkLRPriorities := sets.Int32KeySet(sdkLRByPriority)
	for _, priority := range resLRPriorities.Intersection(sdkLRPriorities).List() {
		resLR := resLRByPriority[priority]
		sdkLR := sdkLRByPriority[priority]
		matchedResAndSDKLRs = append(matchedResAndSDKLRs, resAndSDKListenerRulePair{
			resLR: resLR,
			sdkLR: sdkLR,
		})
	}
	for _, priority := range resLRPriorities.Difference(sdkLRPriorities).List() {
		unmatchedResLRs = append(unmatchedResLRs, resLRByPriority[priority])
	}
	for _, priority := range sdkLRPriorities.Difference(resLRPriorities).List() {
		unmatchedSDKLRs = append(unmatchedSDKLRs, sdkLRByPriority[priority])
	}

	return matchedResAndSDKLRs, unmatchedResLRs, unmatchedSDKLRs
}

func mapResListenerRuleByPriority(resLRs []*elbv2model.ListenerRule) map[int32]*elbv2model.ListenerRule {
	resLRByPriority := make(map[int32]*elbv2model.ListenerRule, len(resLRs))
	for _, resLR := range resLRs {
		resLRByPriority[resLR.Spec.Priority] = resLR
	}
	return resLRByPriority
}

// Filter the rules based on their settings aka actions and conditions
func (s *listenerRuleSynthesizer) matchResAndSDKListenerRulesBySettings(matchedResAndSDKLRsByPriority []resAndSDKListenerRulePair) ([]resAndSDKListenerRulePair, []resAndSDKListenerRulePair, []*resLRDesiredActionsAndConditionsPair, error) {
	var unmatchedResAndSDKListenerRulesBySettings []resAndSDKListenerRulePair
	var matchedResAndSDKListenerRulesBySettings []resAndSDKListenerRulePair
	var unmatchedResLRDesiredActionsAndConditionsPair []*resLRDesiredActionsAndConditionsPair
	for _, resAndSDKLR := range matchedResAndSDKLRsByPriority {
		desiredActions, err := buildSDKActions(resAndSDKLR.resLR.Spec.Actions, s.featureGates)
		if err != nil {
			return nil, nil, nil, err
		}
		desiredConditions := buildSDKRuleConditions(resAndSDKLR.resLR.Spec.Conditions)
		if !cmp.Equal(desiredActions, resAndSDKLR.sdkLR.ListenerRule.Actions, elbv2equality.CompareOptionForActions()) ||
			!cmp.Equal(desiredConditions, resAndSDKLR.sdkLR.ListenerRule.Conditions, elbv2equality.CompareOptionForRuleConditions()) {
			unmatchedResAndSDKListenerRulesBySettings = append(unmatchedResAndSDKListenerRulesBySettings, resAndSDKLR)
			// Storing these so that we don't need to build SDK actions and conditions later during building create request input
			unmatchedResLRDesiredActionsAndConditionsPair = append(unmatchedResLRDesiredActionsAndConditionsPair, &resLRDesiredActionsAndConditionsPair{
				desiredActions:    desiredActions,
				desiredConditions: desiredConditions,
			})
		} else {
			matchedResAndSDKListenerRulesBySettings = append(matchedResAndSDKListenerRulesBySettings, resAndSDKLR)
		}
	}

	return matchedResAndSDKListenerRulesBySettings, unmatchedResAndSDKListenerRulesBySettings, unmatchedResLRDesiredActionsAndConditionsPair, nil
}

func mapSDKListenerRuleByPriority(sdkLRs []ListenerRuleWithTags) map[int32]ListenerRuleWithTags {
	sdkLRByPriority := make(map[int32]ListenerRuleWithTags, len(sdkLRs))
	for _, sdkLR := range sdkLRs {
		priority, _ := strconv.ParseInt(awssdk.ToString(sdkLR.ListenerRule.Priority), 10, 64)
		sdkLRByPriority[int32(priority)] = sdkLR
	}
	return sdkLRByPriority
}

func mapResListenerRuleByListenerARN(resLRs []*elbv2model.ListenerRule) (map[string][]*elbv2model.ListenerRule, error) {
	resLRsByLSARN := make(map[string][]*elbv2model.ListenerRule, len(resLRs))
	ctx := context.Background()
	for _, lr := range resLRs {
		lsARN, err := lr.Spec.ListenerARN.Resolve(ctx)
		if err != nil {
			return nil, err
		}
		resLRsByLSARN[lsARN] = append(resLRsByLSARN[lsARN], lr)
	}
	return resLRsByLSARN, nil
}
