package moleculer

import (
	"context"

	"github.com/moleculer-go/moleculer/broker"
	"github.com/moleculer-go/moleculer/params"
	"github.com/moleculer-go/moleculer/service"
)

type Service = service.ServiceSchema
type Action = service.ServiceActionSchema
type Event = service.ServiceEventSchema
type Params = params.Params

// returns a valid broker based on a passed context
// this is called from any action / event
func BrokerFromContext(ctx *context.Context) *broker.ServiceBroker {
	return broker.FromContext(ctx)
}

// returns a valid broker based on environment configuration
// this is usually called when creating a broker to starting the service(s)
func BrokerFromConfig() *broker.ServiceBroker {
	return broker.FromConfig()
}
