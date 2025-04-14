package interfaces

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/recomendation/entity"
	"github.com/osamikoyo/publics/internal/modules/recomendation/service"
)

type AddFavouriteConntroller struct {
	responder *web.Responder
	service   service.RecomendationService
}

func (conn *AddFavouriteConntroller) Inject(responder *web.Responder, service service.RecomendationService) *AddFavouriteConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *AddFavouriteConntroller) AddFavourite(ctx context.Context, req *web.Request) web.Result {
	var topic entity.Topic

	if !req.QueryAll().Has("id") {
		return conn.responder.BadRequestWithContext(ctx, errors.New("param id not found"))
	}

	id, err := strconv.Atoi(req.QueryAll().Get("id"))
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err = json.NewDecoder(req.Request().Body).Decode(&topic); err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err = conn.service.AddFavTopic(uint(id), topic.ID); err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(map[string]string{
		"message": "success",
	})
}
