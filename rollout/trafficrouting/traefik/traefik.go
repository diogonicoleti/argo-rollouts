package traefik

import (
	"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	logutil "github.com/argoproj/argo-rollouts/utils/log"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	extensionslisters "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/record"
)

// Type holds this controller type
const Type = "Traefik"

// ReconcilerConfig describes static configuration data for the traefik reconciler
type ReconcilerConfig struct {
	Rollout        *v1alpha1.Rollout
	Client         kubernetes.Interface
	Recorder       record.EventRecorder
	ControllerKind schema.GroupVersionKind
	IngressLister  extensionslisters.IngressLister
}

// Reconciler holds required fields to reconcile Traefik resources
type Reconciler struct {
	cfg ReconcilerConfig
	log *logrus.Entry
}

// NewReconciler returns a reconciler struct that brings the canary Ingress into the desired state
func NewReconciler(cfg ReconcilerConfig) *Reconciler {
	return &Reconciler{
		cfg: cfg,
		log: logutil.WithRollout(cfg.Rollout),
	}
}

// Type indicates this reconciler is an Traefik reconciler
func (r *Reconciler) Type() string {
	return Type
}

// Reconcile modifies Traefik Ingress resources to reach desired state
func (r *Reconciler) Reconcile(desiredWeight int32) error {
	return nil
}
