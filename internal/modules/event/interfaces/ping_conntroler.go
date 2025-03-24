package interfaces

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type PingConntroller struct {
	responder *web.Responder
}

func (conn *PingConntroller) Inject(responder *web.Responder) *PingConntroller {
	conn.responder = responder

	return conn
}

func (conn *PingConntroller) Ping(ctx context.Context, r *web.Request) web.Result {
	return conn.responder.Data(map[string]string{
		"hello": "world",
	})
}
