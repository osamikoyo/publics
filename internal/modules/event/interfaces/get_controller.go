package interfaces

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	service "github.com/osamikoyo/publics/internal/modules/event/service"
)

type GetConnroller struct{
	service service.EventService
	responder *web.Responder
}

func (conn *GetConnroller) Inject(responder *web.Responder, service service.EventService) *GetConnroller {
	conn.responder = responder
	conn.service = service

	return conn
}

func (conn *GetConnroller) Get(ctx context.Context, req *web.Request) web.Result {
	id := req.QueryAll().Get("id")
	if id == "" {
		return conn.responder.Data(struct{
			Status int
			Message string
		}{
			Status: http.StatusBadRequest,
			Message: "param id not found",
		})
	}

	event, err := conn.service.GetBy("ID", id)
	if err != nil{
		return conn.responder.Data(struct{
			Status int
			Message string
		}{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return conn.responder.Data(struct{
		Status int
		Message interface{}
	}{
		Status: http.StatusInternalServerError,
		Message: event,
	})
}