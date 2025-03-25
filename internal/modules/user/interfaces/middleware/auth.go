package middleware

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/osamikoyo/publics/internal/modules/user/service"
)

type AuthMiddleware struct {
	service service.UserService
}

func (mw *AuthMiddleware) Inject(service service.UserService) *AuthMiddleware {
	mw.service = service
	return mw
}

func (mw *AuthMiddleware) Auth(ctx context.Context, r *web.Request, w http.ResponseWriter, chain *web.FilterChain) web.Result {
	token := r.Request().Header.Get("Authentification")
	if token == "" {
		return &web.DataResponse{
			Data: struct {
				Message string
				Status  int
			}{
				Message: "token empty",
				Status:  http.StatusNonAuthoritativeInfo,
			},
		}
	}

	claims, err := mw.service.Auth(token)
	if err != nil || claims == nil {
		return &web.DataResponse{
			Data: struct {
				Message string
				Status  int
			}{
				Message: "token empty",
				Status:  http.StatusUnauthorized,
			},
		}
	}

	newCtx := context.WithValue(ctx, "claims", claims)

	return chain.Next(newCtx, r, w)
}
