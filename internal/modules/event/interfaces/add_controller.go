package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
)

type AddConntroller struct{
	service service.EventService
	responder *web.Responder
}

func (conn *AddConntroller) Inject(service service.EventService, responder *web.Responder) *AddConntroller {
	conn.service = service
	conn.responder = responder

	return conn
}

func (a *AddConntroller) Add(ctx context.Context, req *web.Request) web.Result {
	var event entity.Event

	if err := json.NewDecoder(req.Request().Body).Decode(&event); err != nil{
		return a.responder.Data(struct{
			Status int
			Message string
		}{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := a.service.Add(&event);err != nil{
		return a.responder.Data(struct{
			Status int
			Message string
		}{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return a.responder.Data(struct{
		Status int
		Message string
	}{
		Status: http.StatusCreated,
		Message: "success",
	})
}