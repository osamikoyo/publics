package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
)

type UpdateConntroller struct {
	service   service.EventService
	responder *web.Responder
}

func (u *UpdateConntroller) Inject(responder *web.Responder, service service.EventService) *UpdateConntroller {
	u.responder = responder
	u.service = service

	return u
}

func (u *UpdateConntroller) Update(ctx context.Context, req *web.Request) web.Result {
	var updateReq entity.UpdateReq

	if err := json.NewDecoder(req.Request().Body).Decode(&updateReq); err != nil {
		return u.responder.BadRequestWithContext(ctx, err)
	}

	if err := u.service.Update(updateReq.ID, &updateReq.Entity); err != nil {
		return u.responder.ServerErrorWithContext(ctx, err)
	}

	return u.responder.Data(struct {
		Status  int
		Message string
	}{
		Status:  http.StatusOK,
		Message: "success",
	})
}
