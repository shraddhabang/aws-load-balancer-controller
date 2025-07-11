package gateway

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	elbv2gw "sigs.k8s.io/aws-load-balancer-controller/apis/gateway/v1beta1"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/algorithm"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/k8s"
	"sigs.k8s.io/aws-load-balancer-controller/test/framework"
	"sigs.k8s.io/aws-load-balancer-controller/test/framework/utils"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	"strconv"
)

func newCommonResourceStack(dps []*appsv1.Deployment, svcs []*corev1.Service, gwc *gwv1.GatewayClass, gw *gwv1.Gateway, lbc *elbv2gw.LoadBalancerConfiguration, tgcs []*elbv2gw.TargetGroupConfiguration, baseName string, enablePodReadinessGate bool) *commonResourceStack {
	return &commonResourceStack{
		dps:                    dps,
		svcs:                   svcs,
		gwc:                    gwc,
		gw:                     gw,
		lbc:                    lbc,
		tgcs:                   tgcs,
		baseName:               baseName,
		enablePodReadinessGate: enablePodReadinessGate,
	}
}

// commonResourceStack contains resources that are common between nlb / alb gateways
type commonResourceStack struct {
	// configurations
	svcs                   []*corev1.Service
	dps                    []*appsv1.Deployment
	gwc                    *gwv1.GatewayClass
	gw                     *gwv1.Gateway
	lbc                    *elbv2gw.LoadBalancerConfiguration
	tgcs                   []*elbv2gw.TargetGroupConfiguration
	ns                     *corev1.Namespace
	baseName               string
	enablePodReadinessGate bool

	// runtime variables
	createdDPs  []*appsv1.Deployment
	createdSVCs []*corev1.Service
	createdGW   *gwv1.Gateway
}

func (s *commonResourceStack) Deploy(ctx context.Context, f *framework.Framework, resourceSpecificCreation func(ctx context.Context, f *framework.Framework, namespace string) error) error {
	if err := s.allocateNamespace(ctx, f); err != nil {
		return err
	}
	for _, v := range s.dps {
		v.Namespace = s.ns.Name
	}

	for _, v := range s.svcs {
		v.Namespace = s.ns.Name
	}

	for _, v := range s.tgcs {
		v.Namespace = s.ns.Name
	}

	s.gw.Namespace = s.ns.Name
	s.lbc.Namespace = s.ns.Name

	if err := s.createGatewayClass(ctx, f); err != nil {
		return err
	}
	if err := s.createLoadBalancerConfig(ctx, f); err != nil {
		return err
	}
	if err := s.createTargetGroupConfigs(ctx, f); err != nil {
		return err
	}
	if err := s.createDeployments(ctx, f); err != nil {
		return err
	}
	if err := s.createServices(ctx, f); err != nil {
		return err
	}
	if err := s.createGateway(ctx, f); err != nil {
		return err
	}

	if err := resourceSpecificCreation(ctx, f, s.ns.Name); err != nil {
		return err
	}

	if err := s.waitUntilDeploymentReady(ctx, f); err != nil {
		return err
	}

	if err := s.waitUntilServiceReady(ctx, f); err != nil {
		return err
	}

	if err := s.waitUntilGatewayReady(ctx, f); err != nil {
		return err
	}
	return nil
}

func (s *commonResourceStack) Cleanup(ctx context.Context, f *framework.Framework) {
	_ = s.deleteNamespace(ctx, f)
	_ = s.deleteGatewayClass(ctx, f)
}

func (s *commonResourceStack) GetLoadBalancerIngressHostname() string {
	return s.createdGW.Status.Addresses[0].Value
}

func (s *commonResourceStack) getListenersPortMap() map[string]string {
	listenersMap := map[string]string{}
	for _, l := range s.createdGW.Spec.Listeners {
		listenersMap[strconv.Itoa(int(l.Port))] = string(l.Protocol)
	}
	return listenersMap
}

func (s *commonResourceStack) createDeployments(ctx context.Context, f *framework.Framework) error {
	for _, dp := range s.dps {
		f.Logger.Info("creating deployment", "dp", k8s.NamespacedName(dp))
		if err := f.K8sClient.Create(ctx, dp); err != nil {
			f.Logger.Info("failed to create deployment")
			return err
		}
		f.Logger.Info("created deployment", "dp", k8s.NamespacedName(dp))
	}
	return nil
}

func (s *commonResourceStack) waitUntilDeploymentReady(ctx context.Context, f *framework.Framework) error {
	for _, dp := range s.dps {
		f.Logger.Info("waiting until deployment becomes ready", "dp", k8s.NamespacedName(dp))
		_, err := f.DPManager.WaitUntilDeploymentReady(ctx, dp)
		if err != nil {
			f.Logger.Info("failed waiting for deployment")
			return err
		}
		f.Logger.Info("deployment is ready", "dp", k8s.NamespacedName(dp))
	}
	return nil
}

func (s *commonResourceStack) createServices(ctx context.Context, f *framework.Framework) error {
	for _, svc := range s.svcs {
		f.Logger.Info("creating service", "svc", k8s.NamespacedName(svc))
		if err := f.K8sClient.Create(ctx, svc); err != nil {
			f.Logger.Info("failed to create service")
			return err
		}
		f.Logger.Info("created service", "svc", k8s.NamespacedName(svc))
	}
	return nil
}

func (s *commonResourceStack) createGatewayClass(ctx context.Context, f *framework.Framework) error {
	f.Logger.Info("creating gateway class", "gwc", k8s.NamespacedName(s.gwc))
	return f.K8sClient.Create(ctx, s.gwc)
}

func (s *commonResourceStack) createLoadBalancerConfig(ctx context.Context, f *framework.Framework) error {
	f.Logger.Info("creating loadbalancer config", "lbc", k8s.NamespacedName(s.lbc))
	return f.K8sClient.Create(ctx, s.lbc)
}

func (s *commonResourceStack) createTargetGroupConfigs(ctx context.Context, f *framework.Framework) error {
	for _, tgc := range s.tgcs {
		f.Logger.Info("creating target group config", "tgc", k8s.NamespacedName(tgc))
		err := f.K8sClient.Create(ctx, tgc)
		if err != nil {
			f.Logger.Error(err, "failed to create target group config")
			return err
		}
		f.Logger.Info("created target group config", "tgc", k8s.NamespacedName(tgc))
	}
	return nil
}

func (s *commonResourceStack) createGateway(ctx context.Context, f *framework.Framework) error {
	f.Logger.Info("creating gateway", "gw", k8s.NamespacedName(s.gw))
	return f.K8sClient.Create(ctx, s.gw)
}

func (s *commonResourceStack) waitUntilServiceReady(ctx context.Context, f *framework.Framework) error {
	for _, svc := range s.svcs {
		observedSvc := &corev1.Service{}
		err := f.K8sClient.Get(ctx, k8s.NamespacedName(svc), observedSvc)
		if err != nil {
			f.Logger.Error(err, "unable to observe service go ready")
			return err
		}
	}
	return nil
}

func (s *commonResourceStack) waitUntilGatewayReady(ctx context.Context, f *framework.Framework) error {
	observedGw := &gwv1.Gateway{}
	err := wait.PollImmediateUntil(utils.PollIntervalShort, func() (bool, error) {
		if err := f.K8sClient.Get(ctx, k8s.NamespacedName(s.gw), observedGw); err != nil {
			return false, err
		}

		if observedGw.Status.Conditions != nil {
			for _, cond := range observedGw.Status.Conditions {
				if cond.Type == string(gwv1.GatewayConditionProgrammed) && cond.Status == metav1.ConditionTrue {
					return true, nil
				}
			}
		}

		return false, nil
	}, ctx.Done())
	if err != nil {
		return err
	}
	s.createdGW = observedGw
	return nil
}

func (s *commonResourceStack) deleteGatewayClass(ctx context.Context, f *framework.Framework) error {
	return f.K8sClient.Delete(ctx, s.gwc)
}

func (s *commonResourceStack) allocateNamespace(ctx context.Context, f *framework.Framework) error {
	f.Logger.Info("allocating namespace")
	ns, err := f.NSManager.AllocateNamespace(ctx, s.baseName)
	if err != nil {
		return err
	}
	s.ns = ns
	f.Logger.Info("allocated namespace", "nsName", s.ns.Name)
	if s.enablePodReadinessGate {
		f.Logger.Info("label namespace for podReadinessGate injection", "nsName", s.ns.Name)
		oldNS := s.ns.DeepCopy()
		s.ns.Labels = algorithm.MergeStringMap(map[string]string{
			"elbv2.k8s.aws/pod-readiness-gate-inject": "enabled",
		}, s.ns.Labels)
		err := f.K8sClient.Patch(ctx, ns, client.MergeFrom(oldNS))
		if err != nil {
			return err
		}
		f.Logger.Info("labeled namespace with podReadinessGate injection", "nsName", s.ns.Name)
	}
	return nil
}

func (s *commonResourceStack) deleteNamespace(ctx context.Context, tf *framework.Framework) error {
	tf.Logger.Info("deleting namespace", "ns", k8s.NamespacedName(s.ns))
	if err := tf.K8sClient.Delete(ctx, s.ns); err != nil {
		tf.Logger.Info("failed to delete namespace", "ns", k8s.NamespacedName(s.ns))
		return err
	}
	if err := tf.NSManager.WaitUntilNamespaceDeleted(ctx, s.ns); err != nil {
		tf.Logger.Info("failed to wait for namespace deletion", "ns", k8s.NamespacedName(s.ns))
		return err
	}
	tf.Logger.Info("deleted namespace", "ns", k8s.NamespacedName(s.ns))
	return nil
}
