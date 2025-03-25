package user

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces/middleware"
	"github.com/osamikoyo/publics/internal/modules/user/repository"
	"github.com/osamikoyo/publics/internal/modules/user/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModule struct{}

type Routes struct {
	register *interfaces.RegisterConntroller
	login    *interfaces.LoginConntroller
	authMW   *middleware.AuthMiddleware
}

func (r *Routes) Inject(register *interfaces.RegisterConntroller, login *interfaces.LoginConntroller, auth *middleware.AuthMiddleware) *Routes {
	r.register = register
	r.login = login
	r.authMW = auth

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/api/user/register", "user.create")
	registry.MustRoute("/api/user/login", "user.login")

	registry.HandlePost("user.create", r.register.Register)
	registry.HandlePost("user.login", r.login.Login)
}

func (u *UserModule) Configure(inject *dingo.Injector) {
	inject.Bind(new(service.UserService)).To(new(service.UserPrivateService))
	inject.Bind(new(repository.UserRepository)).To(new(repository.UserStorage))
	inject.Bind(new(*gorm.DB)).ToProvider(func() (*gorm.DB, error) {
		return gorm.Open(sqlite.Open("storage/main.db"))
	})

	web.BindRoutes(inject, new(Routes))
}
