package member

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/member/interfaces"
	"github.com/osamikoyo/publics/internal/modules/member/interfaces/middleware"
	"github.com/osamikoyo/publics/internal/modules/member/repository"
	"github.com/osamikoyo/publics/internal/modules/member/service"
)

type MemberModule struct{}

func (m *MemberModule) Inject() *MemberModule {
	return m
}

type Routes struct {
	get    *interfaces.GetConntroller
	add    *interfaces.AddConntroller
	delete *interfaces.DeleteConntroller
	check  middleware.CheckPermsMiddleware
}

func (r *Routes) Inject(get *interfaces.GetConntroller, add *interfaces.AddConntroller,
	delete *interfaces.DeleteConntroller, check middleware.CheckPermsMiddleware) *Routes {
	r.get = get
	r.add = add
	r.delete = delete
	r.check = check

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/api/member/add", "add")
	registry.MustRoute("/api/member/get/:id", "get")
	registry.MustRoute("/api/member/delete/:id", "delete")

	registry.HandlePost("add", r.add.Add)
	registry.HandleGet("get", r.get.Get)
	registry.HandleDelete("delete", r.delete.Delete)
}

func (m *MemberModule) Configure(inject *dingo.Injector) {
	inject.Bind(new(middleware.CheckPermsMiddleware)).To(new(middleware.CheckPermsMW))
	inject.Bind(new(repository.MemberRepository)).To(new(repository.MemberStorage))
	inject.Bind(new(service.MemberService)).To(new(service.MemberServiceImpl))

	web.BindRoutes(inject, new(Routes))
}
