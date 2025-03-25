package interfaces

import (
	"context"
	"encoding/json"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/internal/modules/user/service"
)

type LoginConntroller struct {
	service   service.UserService
	responder *web.Responder
}

func (conn *LoginConntroller) Inject(responder *web.Responder, service service.UserService) *LoginConntroller {
	conn.responder = responder
	conn.service = service
	return conn
}

func (conn *LoginConntroller) Login(ctx context.Context, req *web.Request) web.Result {
	var loginReq entity.LoginRequest

	if err := json.NewDecoder(req.Request().Body).Decode(&loginReq); err != nil {
		return conn.responder.BadRequestWithContext(ctx, err)
	}

	token, err := conn.service.Login(&loginReq)
	if err != nil {
		return conn.responder.ServerError(err)
	}

	return conn.responder.Data(struct {
		Token string
	}{
		Token: token,
	})
}

