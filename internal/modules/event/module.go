package event

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/event/interfaces"
	"github.com/osamikoyo/publics/internal/modules/event/repository"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces/middleware"
  perms_mw 	"github.com/osamikoyo/publics/internal/modules/member/interfaces/middleware"
)

type Config struct {
	CompleteConfig config.Map `inject:"config:user"`
	DSN            string     `inject:"config:user.dsn"`
	Key            string     `inject:"config:user.auth_key"`
}

type EventModule struct {
	cfg *Config
}

func (e *EventModule) Inject(config *Config) *EventModule {
	e.cfg = config

	return e
}

type Routes struct {
	ping   *interfaces.PingConntroller
	add    *interfaces.AddConntroller
	get    *interfaces.GetConnroller
	update *interfaces.UpdateConntroller
	delete *interfaces.DeleteConntroller
	auth   middleware.AuthMW
  checkPerms perms_mw.CheckPermsMiddleware
}

func (r *Routes) Inject(conn *interfaces.PingConntroller, add *interfaces.AddConntroller,
	get *interfaces.GetConnroller, delte *interfaces.DeleteConntroller,
	update *interfaces.UpdateConntroller, auth middleware.AuthMW
  checkPerms perms_mw.CheckPermsMiddleware) *Routes {
	r.ping = conn
	r.add = add
	r.get = get
	r.update = update
	r.delete = delte
	r.auth = auth
  r.checkPerms = checkPerms

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/ping", "ping")
	registry.MustRoute("/api/event/delete/:id", "delete")
	registry.MustRoute("/api/event/update/:id", "update")
	registry.MustRoute("/api/event/create", "create")
	registry.MustRoute("/api/event/get/:id", "get")

	registry.HandleDelete("delete", r.checkPerms.CheckPerms(r.auth.Filter(r.delete.Delete)))
	registry.HandleGet("get", r.get.Get)
	registry.HandlePut("update", r.checkPerms.CheckPerms(r.auth.Filter(r.update.Update)))
	registry.HandlePost("create", r.auth.Filter(r.add.Add))
	registry.HandleGet("ping", r.ping.Ping)
}

type serviceConfig struct {
	Key string
}

func (s *serviceConfig) GetKey() string {
	return s.Key
}

func (e *EventModule) Configure(inject *dingo.Injector) {
  inject.Bind(new(perms_mw.CheckPermsMiddleware)).To(new(perms_mw.CheckPermsMW))
	inject.Bind(new(middleware.AuthMW)).To(new(middleware.AuthMiddleware))

	inject.Bind(new(repository.EventRepository)).To(new(repository.EventStorage))
	inject.Bind(new(service.EventService)).To(new(service.EventServiceImpl))

	web.BindRoutes(inject, new(Routes))
}
