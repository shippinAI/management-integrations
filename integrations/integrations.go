package integrations

import (
	"context"

	"github.com/gorilla/mux"

	"github.com/shippinAI/wardenclyffe/management/internals/modules/peers"
	"github.com/shippinAI/wardenclyffe/util/crypt"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/metric"

	"github.com/shippinAI/wardenclyffe/management/server/store"
	"github.com/shippinAI/wardenclyffe/management/server/telemetry"

	"github.com/shippinAI/wardenclyffe/management/server/account"
	"github.com/shippinAI/wardenclyffe/management/server/activity"
	activitystore "github.com/shippinAI/wardenclyffe/management/server/activity/store"
	"github.com/shippinAI/wardenclyffe/management/server/integrations/integrated_validator"
	"github.com/shippinAI/wardenclyffe/management/server/integrations/port_forwarding"
	"github.com/shippinAI/wardenclyffe/management/server/permissions"
	"github.com/shippinAI/wardenclyffe/management/server/settings"
)

func RegisterHandlers(
	ctx context.Context,
	prefix string,
	router *mux.Router,
	accountManager account.Manager,
	integratedValidator integrated_validator.IntegratedValidator,
	meter metric.Meter,
	permissionsManager permissions.Manager,
	peersManager peers.Manager,
	proxyController port_forwarding.Controller,
	settingsManager settings.Manager,
) (*mux.Router, error) {
	return router, nil
}

func InitEventStore(ctx context.Context, dataDir string, key string, _ *Metrics) (activity.Store, string, error) {
	var err error
	if key == "" {
		log.Debugf("generate new activity store encryption key")
		key, err = crypt.GenerateKey()
		if err != nil {
			return nil, "", err
		}
	}
	store, err := activitystore.NewSqlStore(ctx, dataDir, key)
	return store, key, err
}

func InitPermissionsManager(store store.Store, metric metric.Meter) permissions.Manager {
	return permissions.NewManager(store)
}

type Metrics struct {
	telemetry.AppMetrics
}

func InitIntegrationMetrics(ctx context.Context, metrics telemetry.AppMetrics) (*Metrics, error) {
	return &Metrics{
		AppMetrics: metrics,
	}, nil
}

func IsValidChildAccount(_ context.Context, _, _, _ string) bool {
	return false
}
