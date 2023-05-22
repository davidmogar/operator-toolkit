package webhook

import (
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Webhook is an interface that should be implemented by operator webhooks and that allows to have a cohesive way
// of defining them and register them.
type Webhook interface {
	Register(mgr ctrl.Manager, log *logr.Logger) error
}

// SetupWebhooks invoke the Register function of every webhook passed as an argument to this function.
func SetupWebhooks(mgr manager.Manager, webhooks ...Webhook) error {
	log := ctrl.Log.WithName("webhooks")

	for _, webhook := range webhooks {
		err := webhook.Register(mgr, &log)
		if err != nil {
			return err
		}
	}

	return nil
}
