package middleware

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/golang-jwt/jwt/v5"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces/config"
)

type AuthMW interface {
	Auth(ctx context.Context, r *web.Request, w *http.ResponseWriter, chain *web.FilterChain) web.Result
}

type Claims struct {
	jwt.RegisteredClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type AuthMiddleware struct {
	key config.ServiceConfig
}

func (mw *AuthMiddleware) Inject() *AuthMiddleware {
	return mw
}

func (mw *AuthMiddleware) Auth(ctx context.Context, r *web.Request, w http.ResponseWriter, chain *web.FilterChain) web.Result {
	tkn := r.Request().Header.Get("Authentification")
	if tkn == "" {
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

	token, err := jwt.ParseWithClaims(tkn, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mw.key.GetKey()), nil
	})

	if err != nil {
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

	claims, ok := token.Claims.(*Claims)
	if !ok {
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
