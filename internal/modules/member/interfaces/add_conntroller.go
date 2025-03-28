package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/member/entity"
	"github.com/osamikoyo/publics/internal/modules/member/service"
)

type AddConntroller struct {
	service   service.MemberService
	responder *web.Responder
}

func (conn *AddConntroller) Inject(service service.MemberService, responder *web.Responder) *AddConntroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *AddConntroller) Add(ctx context.Context, req *web.Request) web.Result {
	var member entity.PublicMember

	if err := json.NewDecoder(req.Request().Body).Decode(&member); err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	if err := conn.service.Add(&member); err != nil {
		return conn.responder.ServerErrorWithContext(ctx, err)
	}

	return conn.responder.Data(
		struct {
			Status  int
			Message string
		}{
			Status:  http.StatusCreated,
			Message: "succes",
		},
	)
}
