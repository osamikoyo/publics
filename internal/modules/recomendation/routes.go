package recomendation

import (
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/recomendation/interfaces"
)

type Routes struct {
	addFav *interfaces.AddFavouriteConntroller
	getRec *interfaces.GetRecsConntroller
}

func (r *Routes) Inject(addFav *interfaces.AddFavouriteConntroller, getRec *interfaces.GetRecsConntroller) *Routes {
	r.addFav = addFav
	r.getRec = getRec

	return r
}

func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/recomendation/add", "add")
	registry.MustRoute("/recomendation/get", "get")

	registry.HandleGet("get", r.getRec.GetRecs)
	registry.HandlePost("add", r.addFav.AddFavourite)
}
