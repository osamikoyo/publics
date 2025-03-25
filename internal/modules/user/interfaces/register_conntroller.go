package interfaces

import (
	"context"
	"encoding/json"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/internal/modules/user/service"
)

type RegisterConntroller struct{
	service service.UserService
	responder *web.Responder
}

func (conn *RegisterConntroller) Inject(responder *web.Responder) *RegisterConntroller {
	conn.responder = responder
	return conn
}

func (conn *RegisterConntroller) Register(ctx context.Context, r *web.Request) web.Result {
	var user entity.User

	if err := json.NewDecoder(r.Request().Body).Decode(&user);err != nil{
		return conn.responder.BadRequestWithContext(context.Background(), err)
	}

	if err := conn.service.Register(&user);err != nil{
		return conn.responder.ServerError(err)
	}

	return conn.responder.Data(map[string]string{
		"message" : "success",
	})
}