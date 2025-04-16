package recomendation

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/recomendation/repository"
	"github.com/osamikoyo/publics/internal/modules/recomendation/service"
)

type RecomendationModule struct{}

func (r *RecomendationModule) Inject() *RecomendationModule {
	return r
}

func (r *RecomendationModule) Configure(inject *dingo.Injector) {
	inject.Bind(new(repository.RecomendationRepository)).To(new(repository.RecomendationStorage))

	inject.Bind(new(service.RecomendationService)).To(new(service.RecomendationServiceImpl))

	web.BindRoutes(inject, new(Routes))
}
