package interfaces

import (
	"context"
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
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err := conn.service.Delete(uint(id)); err != nil {
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
