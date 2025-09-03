package k8s

import (
	"testing"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func Test_defaultSecretsManager_MonitorSecrets(t *testing.T) {
	type monitorSecretsCall struct {
		consumerID string
		secrets    []types.NamespacedName
	}
	tests := []struct {
		testName           string
		monitorSecretsCall []monitorSecretsCall
		wantSecrets        []types.NamespacedName
	}{
		{
			testName: "No secrets",
		},
		{
			testName: "Single Ingress consumer",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "ig-group1",
					secrets: []types.NamespacedName{
						{Name: "secret-1", Namespace: "ns-1"},
					},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "secret-1", Namespace: "ns-1"},
			},
		},
		{
			testName: "Single ingress consumer, multiple secrets",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "ig-group1",
					secrets: []types.NamespacedName{
						{Name: "secret-1", Namespace: "ns-1"},
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
					},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "secret-1", Namespace: "ns-1"},
				{Name: "secret-2", Namespace: "ns-2"},
				{Name: "secret-3", Namespace: "ns-3"},
			},
		},
		{
			testName: "Multiple consumers, mix of ingress and gateway, overlapping secrets",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "ig-group1",
					secrets: []types.NamespacedName{
						{Name: "secret-1", Namespace: "ns-1"},
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
					},
				},
				{
					consumerID: "gw-1",
					secrets: []types.NamespacedName{
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
						{Name: "secret-4", Namespace: "ns-4"},
					},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "secret-1", Namespace: "ns-1"},
				{Name: "secret-2", Namespace: "ns-2"},
				{Name: "secret-3", Namespace: "ns-3"},
				{Name: "secret-4", Namespace: "ns-4"},
			},
		},
		{
			testName: "Multiple ingress consumers, with deletion",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "group1",
					secrets: []types.NamespacedName{
						{Name: "secret-1", Namespace: "ns-1"},
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
					},
				},
				{
					consumerID: "group2",
					secrets: []types.NamespacedName{
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
						{Name: "secret-4", Namespace: "ns-4"},
					},
				},
				{
					consumerID: "group1",
					secrets: []types.NamespacedName{
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
					},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "secret-2", Namespace: "ns-2"},
				{Name: "secret-3", Namespace: "ns-3"},
				{Name: "secret-4", Namespace: "ns-4"},
			},
		},
		{
			testName: "Multiple ingress consumers, delete all",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "group1",
					secrets: []types.NamespacedName{
						{Name: "secret-1", Namespace: "ns-1"},
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
					},
				},
				{
					consumerID: "group2",
					secrets: []types.NamespacedName{
						{Name: "secret-2", Namespace: "ns-2"},
						{Name: "secret-3", Namespace: "ns-3"},
						{Name: "secret-4", Namespace: "ns-4"},
					},
				},
				{
					consumerID: "group1",
					secrets:    []types.NamespacedName{},
				},
				{
					consumerID: "group2",
					secrets:    []types.NamespacedName{},
				},
			},
			wantSecrets: []types.NamespacedName{},
		},
		{
			testName: "multiple gateways, overlapping secrets",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "gw-1",
					secrets: []types.NamespacedName{
						{Name: "oidc-secret-1", Namespace: "auth-ns"},
						{Name: "shared-secret", Namespace: "shared-ns"},
					},
				},
				{
					consumerID: "gw-2",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "shared-ns"},
						{Name: "oidc-secret-2", Namespace: "auth-ns"},
					},
				},
				{
					consumerID: "gw-3",
					secrets: []types.NamespacedName{
						{Name: "prod-secret", Namespace: "production"},
						{Name: "shared-secret", Namespace: "shared-ns"},
					},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "oidc-secret-1", Namespace: "auth-ns"},
				{Name: "shared-secret", Namespace: "shared-ns"},
				{Name: "oidc-secret-2", Namespace: "auth-ns"},
				{Name: "prod-secret", Namespace: "production"},
			},
		},
		{
			testName: "multiple gateways, with deletion",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "gw-1",
					secrets: []types.NamespacedName{
						{Name: "oidc-secret-1", Namespace: "auth-ns"},
						{Name: "shared-secret", Namespace: "shared-ns"},
					},
				},
				{
					consumerID: "gw-2",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "shared-ns"},
						{Name: "oidc-secret-2", Namespace: "auth-ns"},
					},
				},
				{
					consumerID: "gw-3",
					secrets: []types.NamespacedName{
						{Name: "prod-secret", Namespace: "production"},
						{Name: "shared-secret", Namespace: "shared-ns"},
					},
				},
				// Delete the first gateway configuration
				{
					consumerID: "gw-1",
					secrets:    []types.NamespacedName{},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "shared-secret", Namespace: "shared-ns"}, // Still used by gw-2 and gw-3
				{Name: "oidc-secret-2", Namespace: "auth-ns"},   // Still used by gw-2
				{Name: "prod-secret", Namespace: "production"},  // Still used by gw-3
			},
		},
		{
			testName: "Cross-controller cleanup - ingress removed, multiple gateways remain",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "ig-group1",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "ingress-only-secret", Namespace: "ns-1"},
					},
				},
				{
					consumerID: "gw-1",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "gateway-secret-1", Namespace: "ns-2"},
					},
				},
				{
					consumerID: "gw-2",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "gateway-secret-2", Namespace: "ns-3"},
					},
				},
				// Remove ingress consumer
				{
					consumerID: "ig-group1",
					secrets:    []types.NamespacedName{},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "shared-secret", Namespace: "ns-1"},    // Should still exist for both gateways
				{Name: "gateway-secret-1", Namespace: "ns-2"}, // Should still exist for gw 1
				{Name: "gateway-secret-2", Namespace: "ns-3"}, // Should still exist for gw 2
			},
		},
		{
			testName: "Cross-controller cleanup - one gateway removed, ingress and other gateway remain",
			monitorSecretsCall: []monitorSecretsCall{
				{
					consumerID: "ig-group1",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "ingress-only-secret", Namespace: "ns-1"},
					},
				},
				{
					consumerID: "gw-1",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "gateway-secret-1", Namespace: "ns-2"},
					},
				},
				{
					consumerID: "gw-2",
					secrets: []types.NamespacedName{
						{Name: "shared-secret", Namespace: "ns-1"},
						{Name: "gateway-secret-2", Namespace: "ns-3"},
					},
				},
				// Remove one gateway consumer
				{
					consumerID: "gw-1",
					secrets:    []types.NamespacedName{},
				},
			},
			wantSecrets: []types.NamespacedName{
				{Name: "shared-secret", Namespace: "ns-1"},       // Should still exist for ingress and gw 2
				{Name: "ingress-only-secret", Namespace: "ns-1"}, // Should still exist for ingress
				{Name: "gateway-secret-2", Namespace: "ns-3"},    // Should still exist for gw 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			secretsEventChan := make(chan event.TypedGenericEvent[*corev1.Secret], 100)
			fakeClient := fake.NewSimpleClientset()
			secretsManager := NewSecretsManager(fakeClient, secretsEventChan, logr.New(&log.NullLogSink{}))

			for _, call := range tt.monitorSecretsCall {
				secretsManager.MonitorSecrets(call.consumerID, call.secrets)
			}
			assert.Equal(t, len(tt.wantSecrets), len(secretsManager.secretMap))
			for _, want := range tt.wantSecrets {
				_, exists := secretsManager.secretMap[want]
				assert.True(t, exists)
			}
		})
	}
}
