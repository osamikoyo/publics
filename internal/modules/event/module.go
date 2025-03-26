package event

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	"github.com/osamikoyo/publics/internal/modules/event/interfaces"
	"github.com/osamikoyo/publics/internal/modules/event/repository"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
	cfg "github.com/osamikoyo/publics/internal/modules/user/interfaces/config"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	CompleteConfig config.Map `inject:"config:user"`
	DSN            string     `inject:"config:user.dsn"`
	Key            string     `inject:"config:user.auth_key"`
}

type EventModule struct {
	cfg     *Config
	service service.EventService
}

func (e *EventModule) Inject(config *Config, service service.EventService) *EventModule {
	e.cfg = config
	e.service = service

	return e
}

type Routes struct {
	ping   *interfaces.PingConntroller
	add    *interfaces.AddConntroller
	get    *interfaces.GetConnroller
	delete *interfaces.UpdateConntroller
}

func (r *Routes) Inject(conn *interfaces.PingConntroller) *Routes {
	r.ping = conn

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/ping", "ping")

	registry.HandleGet("ping", r.ping.Ping)
}

type serviceConfig struct {
	Key string
}

func (s *serviceConfig) GetKey() string {
	return s.Key
}

func (e *EventModule) Configure(inject *dingo.Injector) {
	inject.Bind(new(cfg.ServiceConfig)).ToProvider(func() cfg.ServiceConfig {
		return &serviceConfig{
			Key: e.cfg.Key,
		}
	})

	inject.Bind((*gorm.DB)(nil)).ToProvider(func() (*gorm.DB, error) {
		db, err := gorm.Open(sqlite.Open(e.cfg.DSN))
		if err != nil {
			return nil, err
		}
		inject.Bind((*logger.Logger)(nil)).ToProvider(func() *logger.Logger {
			return logger.Init()
		})

		return db, db.AutoMigrate(&entity.Event{})
	})

	inject.Bind(new(repository.EventRepository)).To(new(repository.EventStorage))
	inject.Bind(new(service.EventService)).To(new(service.EventServiceImpl))

	web.BindRoutes(inject, new(Routes))
}
