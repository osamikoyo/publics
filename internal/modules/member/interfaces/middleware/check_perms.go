package middleware

import (
	"context"
	"net/http"
	"strconv"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/golang-jwt/jwt/v5"
	"github.com/osamikoyo/publics/internal/modules/member/service"
)

type Claims struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type CheckPermsMiddleware interface {
	CheckPerms(web.Action) web.Action
}

type CheckPermsMW struct {
	service service.MemberService
}

func (mw *CheckPermsMW) Inject(service service.MemberService) *CheckPermsMW {
	mw.service = service
	return mw
}

func (mw *CheckPermsMW) CheckPerms(next web.Action) web.Action {
	return func(ctx context.Context, req *web.Request) web.Result {
		claims, ok := ctx.Value("claims").(*Claims)
		if !ok {
			return &web.Response{
				Status: http.StatusUnauthorized,
			}
		}

		strID := req.Params["id"]
		if strID == "" {
			return &web.Response{
				Status: http.StatusBadRequest,
			}
		}

		id, err := strconv.Atoi(strID)
		if err != nil {
			return &web.ServerErrorResponse{
				Error: err,
			}
		}

		ok, err = mw.service.CheckPerms(claims.ID, uint(id))
		if err != nil {
			return &web.ServerErrorResponse{
				Error: err,
			}
		}

		if !ok || err != nil {
			return &web.Response{
				Status: http.StatusUnauthorized,
			}
		} else {
			return next(ctx, req)
		}
	}
}
