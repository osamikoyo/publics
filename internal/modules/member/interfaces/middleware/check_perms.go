package middleware

import "flamingo.me/flamingo/v3/framework/web"

type CheckPermsMiddleware interface {
	CheckPerms(web.Action) web.Action
}

type CheckPermsMW struct {
}
