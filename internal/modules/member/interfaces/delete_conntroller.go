package interfaces

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/member/service"
)

type DeleteConntroller struct {
	service   service.MemberService
	responder *web.Responder
}

func (conn *DeleteConntroller) Inject(service service.MemberService, responder *web.Responder) *DeleteConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *DeleteConntroller) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return conn.responder.BadRequestWithContext(ctx, errors.New("params id not found"))
	}

	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err = conn.service.Delete(uint(id)); err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(struct {
		Status  int
		Message string
	}{
		Status:  http.StatusOK,
		Message: "success",
	})
}
