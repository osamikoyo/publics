package interfaces

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
)

type DeleteConntroller struct {
	service   service.EventService
	responder *web.Responder
}

func (conn *DeleteConntroller) Inject(responder *web.Responder, service service.EventService) *DeleteConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *DeleteConntroller) Delete(ctx context.Context, req *web.Request) web.Result {
	id := req.QueryAll().Get("id")

	if id == "" {
		return conn.responder.BadRequestWithContext(ctx, errors.New("request without id"))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err := conn.service.Delete(uint(idInt)); err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(
		struct {
			Status  int
			Message string
		}{
			Status:  http.StatusOK,
			Message: "success",
		},
	)
}
