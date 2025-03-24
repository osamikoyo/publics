package event

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/event/interfaces"
)

type EventModule struct {
}

type Routes struct {
	ping *interfaces.PingConntroller
}

func (r *Routes) Inject(conn *interfaces.PingConntroller) *Routes {
	r.ping = conn

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/ping", "ping")

	registry.HandleGet("ping", r.ping.Ping)
}

func (e *EventModule) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
}
