package interfaces

import (
	"context"
	"errors"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/recomendation/service"
)

type GetRecsConntroller struct {
	responder *web.Responder
	service   service.RecomendationService
}

func (conn *GetRecsConntroller) Inject(responder *web.Responder, service service.RecomendationService) *GetRecsConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *GetRecsConntroller) GetRecs(ctx context.Context, req *web.Request) web.Result {
	if !req.QueryAll().Has("id") {
		return conn.responder.BadRequestWithContext(ctx, errors.New("query param id not exist"))
	}

	id, err := strconv.Atoi(req.QueryAll().Get("id"))
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	recs, err := conn.service.GetRecs(uint(id))
	if err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(recs)
}
