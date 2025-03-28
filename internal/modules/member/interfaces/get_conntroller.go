package interfaces

import (
	"context"
	"errors"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/member/service"
)

type GetConntroller struct {
	service   service.MemberService
	responder *web.Responder
}

func (conn *GetConntroller) Inject(service service.MemberService, responder *web.Responder) *GetConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *GetConntroller) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return conn.responder.BadRequestWithContext(ctx, errors.New("param id not found"))
	}

	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	members, err := conn.service.Get(uint(id))
	if err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(members)
}
