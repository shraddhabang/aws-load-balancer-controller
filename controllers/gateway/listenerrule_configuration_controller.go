package gateway

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	elbv2gw "sigs.k8s.io/aws-load-balancer-controller/apis/gateway/v1beta1"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/config"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/gateway/constants"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/gateway/routeutils"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/k8s"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/runtime"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/shared_constants"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

const (
	secretValidationRequeueInterval = 30 * time.Second
)

// NewListenerRuleConfigurationReconciler constructs a reconciler that responds to listener rule configuration changes
func NewListenerRuleConfigurationReconciler(k8sClient client.Client, eventRecorder record.EventRecorder, controllerConfig config.ControllerConfig, finalizerManager k8s.FinalizerManager, logger logr.Logger) Reconciler {
	return &listenerRuleConfigurationReconciler{
		k8sClient:        k8sClient,
		eventRecorder:    eventRecorder,
		logger:           logger,
		finalizerManager: finalizerManager,
		workers:          controllerConfig.GatewayClassMaxConcurrentReconciles,
	}
}

// listenerRuleConfigurationReconciler reconciles listener rule configurations
type listenerRuleConfigurationReconciler struct {
	k8sClient        client.Client
	logger           logr.Logger
	eventRecorder    record.EventRecorder
	secretsManager   k8s.SecretsManager
	finalizerManager k8s.FinalizerManager
	workers          int
}

func (r *listenerRuleConfigurationReconciler) SetupWatches(_ context.Context, ctrl controller.Controller, mgr ctrl.Manager, clientSet *kubernetes.Clientset) error {

	if err := ctrl.Watch(source.Kind(mgr.GetCache(), &elbv2gw.ListenerRuleConfiguration{}, &handler.TypedEnqueueRequestForObject[*elbv2gw.ListenerRuleConfiguration]{})); err != nil {
		return err
	}
	secretEventsChan := make(chan event.TypedGenericEvent[*corev1.Secret])
	secretToLRCHandler := &secretToListenerRuleConfigHandler{
		k8sClient: r.k8sClient,
		logger:    r.logger,
	}
	if err := ctrl.Watch(source.Channel(secretEventsChan, secretToLRCHandler)); err != nil {
		return err
	}
	r.secretsManager = k8s.NewSecretsManager(clientSet, secretEventsChan, r.logger.WithName("secrets-manager"))
	return nil
}

func (r *listenerRuleConfigurationReconciler) Reconcile(ctx context.Context, req reconcile.Request) (ctrl.Result, error) {
	return runtime.HandleReconcileError(r.reconcile(ctx, req), r.logger)
}

func (r *listenerRuleConfigurationReconciler) reconcile(ctx context.Context, req reconcile.Request) error {
	listenerRuleConf := &elbv2gw.ListenerRuleConfiguration{}
	if err := r.k8sClient.Get(ctx, req.NamespacedName, listenerRuleConf); err != nil {
		return client.IgnoreNotFound(err)
	}

	r.logger.V(1).Info("Reconcile request for listener rule configuration", "cfg", listenerRuleConf)

	if listenerRuleConf.DeletionTimestamp == nil || listenerRuleConf.DeletionTimestamp.IsZero() {
		return r.handleUpdate(ctx, listenerRuleConf)
	}

	return r.handleDelete(listenerRuleConf)
}

func (r *listenerRuleConfigurationReconciler) handleUpdate(ctx context.Context, listenerRuleConf *elbv2gw.ListenerRuleConfiguration) error {
	if !k8s.HasFinalizer(listenerRuleConf, shared_constants.ListenerRuleConfigurationFinalizer) {
		if err := r.finalizerManager.AddFinalizers(context.Background(), listenerRuleConf, shared_constants.ListenerRuleConfigurationFinalizer); err != nil {
			return err
		}
	}
	inUse, err := routeutils.IsListenerRuleConfigInUse(context.Background(), listenerRuleConf, r.k8sClient)
	if err != nil {
		return fmt.Errorf("skipping finalizer removal due failure to verify if listener rule configuration [%+v] is in use. Error : %w ", k8s.NamespacedName(listenerRuleConf), err)
	}
	if inUse {
		secret, secretValidationErr := r.validateRequiredSecrets(ctx, listenerRuleConf)

		if secretValidationErr != nil {
			if isSecretNotFoundError(secretValidationErr) {
				// Update status: NOT accepted
				if err := r.updateStatus(ctx, listenerRuleConf, false, secretValidationErr.Error()); err != nil {
					return err
				}
				return runtime.NewRequeueNeededAfter("Required secret not yet available", secretValidationRequeueInterval)
			}
			return secretValidationErr
		}
		// Update status: ACCEPTED
		if err := r.updateStatus(ctx, listenerRuleConf, true, "Accepted"); err != nil {
			return err
		}
		r.secretsManager.MonitorSecrets(k8s.NamespacedName(listenerRuleConf).String(), secret)
	}
	return nil
}

func (r *listenerRuleConfigurationReconciler) handleDelete(listenerRuleConf *elbv2gw.ListenerRuleConfiguration) error {
	if !k8s.HasFinalizer(listenerRuleConf, shared_constants.ListenerRuleConfigurationFinalizer) {
		return nil
	}

	inUse, err := routeutils.IsListenerRuleConfigInUse(context.Background(), listenerRuleConf, r.k8sClient)

	if err != nil {
		return fmt.Errorf("skipping finalizer removal due failure to verify if listener rule configuration [%+v] is in use. Error : %w ", k8s.NamespacedName(listenerRuleConf), err)
	}
	// if the listener rule configuration is still in use, we should not delete it
	if inUse {
		return fmt.Errorf("failed to remove finalizers as listener rule configuration [%+v] is still in use", k8s.NamespacedName(listenerRuleConf))
	}
	r.secretsManager.MonitorSecrets(k8s.NamespacedName(listenerRuleConf).String(), nil)
	return r.finalizerManager.RemoveFinalizers(context.Background(), listenerRuleConf, shared_constants.ListenerRuleConfigurationFinalizer)
}

func (r *listenerRuleConfigurationReconciler) SetupWithManager(_ context.Context, mgr ctrl.Manager) (controller.Controller, error) {
	return controller.New(constants.ListenerRuleConfigurationController, mgr, controller.Options{
		MaxConcurrentReconciles: r.workers,
		Reconciler:              r,
	})

}

type SecretNotFoundError struct {
	SecretName      string
	SecretNamespace string
}

func (e *SecretNotFoundError) Error() string {
	return fmt.Sprintf("secret [%s/%s] not found", e.SecretNamespace, e.SecretName)
}

func isSecretNotFoundError(err error) bool {
	_, ok := err.(*SecretNotFoundError)
	return ok
}

func (r *listenerRuleConfigurationReconciler) validateRequiredSecrets(ctx context.Context, listenerRuleConf *elbv2gw.ListenerRuleConfiguration) ([]types.NamespacedName, error) {
	if listenerRuleConf.Spec.Actions == nil {
		return nil, nil
	}

	for _, action := range listenerRuleConf.Spec.Actions {
		if action.Type == elbv2gw.ActionTypeAuthenticateOIDC && action.AuthenticateOIDCConfig != nil && action.AuthenticateOIDCConfig.Secret != nil {
			secretName := action.AuthenticateOIDCConfig.Secret.Name
			secretNamespace := listenerRuleConf.Namespace

			if action.AuthenticateOIDCConfig.Secret.Namespace != nil {
				secretNamespace = *action.AuthenticateOIDCConfig.Secret.Namespace
			}
			secret := &corev1.Secret{}
			secretKey := types.NamespacedName{
				Name:      secretName,
				Namespace: secretNamespace,
			}
			err := r.k8sClient.Get(ctx, secretKey, secret)

			if apierrors.IsNotFound(err) {
				return nil, &SecretNotFoundError{
					SecretName:      secretName,
					SecretNamespace: secretNamespace,
				}
			}
			if err != nil {
				return nil, err
			}
			return []types.NamespacedName{secretKey}, nil
		}
	}
	return nil, nil
}

// updateStatus updates the ListenerRuleConfiguration status with the current validation state
func (r *listenerRuleConfigurationReconciler) updateStatus(ctx context.Context, listenerRuleConf *elbv2gw.ListenerRuleConfiguration, accepted bool, message string) error {
	// Check if status actually needs updating
	if listenerRuleConf.Status.Accepted != nil &&
		*listenerRuleConf.Status.Accepted == accepted &&
		listenerRuleConf.Status.Message != nil &&
		*listenerRuleConf.Status.Message == message &&
		listenerRuleConf.Status.ObservedGeneration != nil &&
		*listenerRuleConf.Status.ObservedGeneration == listenerRuleConf.Generation {
		return nil // No update needed
	}

	listenerRuleConfOld := listenerRuleConf.DeepCopy()

	// Update status fields
	listenerRuleConf.Status.Accepted = &accepted
	listenerRuleConf.Status.Message = &message
	listenerRuleConf.Status.ObservedGeneration = &listenerRuleConf.Generation

	// Patch the status
	if err := r.k8sClient.Status().Patch(ctx, listenerRuleConf, client.MergeFrom(listenerRuleConfOld)); err != nil {
		return fmt.Errorf("failed to update ListenerRuleConfiguration status: %w", err)
	}

	return nil
}

// Add this to the controller file
type secretToListenerRuleConfigHandler struct {
	k8sClient client.Client
	logger    logr.Logger
}

func (h *secretToListenerRuleConfigHandler) Create(ctx context.Context, e event.TypedCreateEvent[*corev1.Secret], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	// No-op - predicate filters this out
}

func (h *secretToListenerRuleConfigHandler) Update(ctx context.Context, e event.TypedUpdateEvent[*corev1.Secret], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	//secretOld := e.ObjectOld
	//secretNew := e.ObjectNew
	//
	//// we only care below update event:
	////	1. Secret data updates
	////	2. Secret deletions
	//if equality.Semantic.DeepEqual(secretOld.Data, secretNew.Data) &&
	//	equality.Semantic.DeepEqual(secretOld.DeletionTimestamp.IsZero(), secretNew.DeletionTimestamp.IsZero()) {
	//	return
	//}
	//h.logger.V(1).Info("enqueue secret update event", "secret", secretNew.Name)
	//h.enqueueImpactedListenerRulesConfigs(ctx, secretNew)
}

func (h *secretToListenerRuleConfigHandler) Delete(ctx context.Context, e event.TypedDeleteEvent[*corev1.Secret], queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	secretOld := e.Object
	h.logger.V(1).Info("checking affected ListenerRuleConfigurations", "secret", k8s.NamespacedName(secretOld))

	// 🚀 FIND affected ListenerRuleConfigurations
	affectedLRCs, err := routeutils.FilterListenerRuleConfigBySecret(ctx, h.k8sClient, secretOld)
	if err != nil {
		h.logger.Error(err, "failed to find ListenerRuleConfigurations affected by secret deletion", "secret", k8s.NamespacedName(secretOld))
		return
	}

	// 🚀 ENQUEUE each affected LRC for reconciliation (not the secret!)
	for _, lrc := range affectedLRCs {
		h.logger.V(1).Info("enqueuing ListenerRuleConfiguration for secret deletion", "lrc", k8s.NamespacedName(lrc), "secret", k8s.NamespacedName(secretOld))
		queue.Add(reconcile.Request{
			NamespacedName: k8s.NamespacedName(lrc), // ✅ Enqueue the LRC, not the secret!
		})
	}
}

func (h *secretToListenerRuleConfigHandler) Generic(ctx context.Context, e event.TypedGenericEvent[*corev1.Secret], queue workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	secretOld := e.Object
	h.logger.V(1).Info("checking affected ListenerRuleConfigurations", "secret", k8s.NamespacedName(secretOld))

	// 🚀 FIND affected ListenerRuleConfigurations
	affectedLRCs, err := routeutils.FilterListenerRuleConfigBySecret(ctx, h.k8sClient, secretOld)
	if err != nil {
		h.logger.Error(err, "failed to find ListenerRuleConfigurations affected by secret deletion", "secret", k8s.NamespacedName(secretOld))
		return
	}

	// 🚀 ENQUEUE each affected LRC for reconciliation (not the secret!)
	for _, lrc := range affectedLRCs {
		h.logger.V(1).Info("enqueuing ListenerRuleConfiguration for secret deletion", "lrc", k8s.NamespacedName(lrc), "secret", k8s.NamespacedName(secretOld))
		queue.Add(reconcile.Request{
			NamespacedName: k8s.NamespacedName(lrc), // ✅ Enqueue the LRC, not the secret!
		})
	}
}
