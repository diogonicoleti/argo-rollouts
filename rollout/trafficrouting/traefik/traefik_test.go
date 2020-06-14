package traefik

import (
	"testing"

	"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fakeRollout(stableSvc, canarySvc, stableIng string) *v1alpha1.Rollout {
	return &v1alpha1.Rollout{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "rollout",
			Namespace: metav1.NamespaceDefault,
		},
		Spec: v1alpha1.RolloutSpec{
			Strategy: v1alpha1.RolloutStrategy{
				Canary: &v1alpha1.CanaryStrategy{
					StableService: stableSvc,
					CanaryService: canarySvc,
					TrafficRouting: &v1alpha1.RolloutTrafficRouting{
						Traefik: &v1alpha1.TraefikTrafficRouting{
							Ingress: stableIng,
						},
					},
				},
			},
		},
	}
}

func TestType(t *testing.T) {
	rollout := fakeRollout("stable-service", "canary-service", "stable-ingress")
	r := NewReconciler(ReconcilerConfig{
		Rollout: rollout,
	})
	assert.Equal(t, Type, r.Type())
}
