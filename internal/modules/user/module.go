package user

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces"
	"github.com/osamikoyo/publics/internal/modules/user/repository"
	"github.com/osamikoyo/publics/internal/modules/user/service"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct{
	CompleteConfig config.Map `inject:"config:user"`
	DSN string `inject:"config:user.dsn"`
	Key string `inject:"config:user.key"`
}

type UserModule struct{
	cfg *Config 
}

func (u *UserModule) Inject(cfg *Config) *UserModule {
	u.cfg = cfg
	return u
}

type Routes struct {
	register *interfaces.RegisterConntroller
	login    *interfaces.LoginConntroller
}

func (r *Routes) Inject(register *interfaces.RegisterConntroller, login *interfaces.LoginConntroller) *Routes {
	r.register = register
	r.login = login

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/api/user/register", "user.create")
	registry.MustRoute("/api/user/login", "user.login")

	registry.HandlePost("user.create", r.register.Register)
	registry.HandlePost("user.login", r.login.Login)
}

func (u *UserModule) Configure(inject *dingo.Injector) {
	inject.Bind((*gorm.DB)(nil)).ToProvider(func() (*gorm.DB, error) {
		db, err := gorm.Open(sqlite.Open(u.cfg.DSN))
		if err != nil {
			return nil, err
		}
	inject.Bind((*logger.Logger)(nil)).ToProvider(func () *logger.Logger {
		return logger.Init()
	})

		return db, db.AutoMigrate(&entity.User{})
	})

	inject.Bind(new(repository.UserRepository)).To(new(repository.UserStorage))
	inject.Bind(new(service.UserService)).To(new(service.UserPrivateService))

	web.BindRoutes(inject, new(Routes))
}
