package api

import (
	"ohp/internal/api/handler"
	middle "ohp/internal/api/middleware"
	"ohp/internal/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

func NewRouter(
	pushHandler *handler.PushHandler,
	authHandler *handler.AuthHandler,
	env config.Env,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middle.CorsMiddleware(env.FrontUrl))

	r.Mount("/push", pushHandler.Routes())
	r.Mount("/auth", authHandler.Routes())

	return r
}

var routeModule = fx.Module("router",
	fx.Provide(
		handler.NewPushHandler,
		handler.NewAuthHandler,
	),

	fx.Provide(NewRouter),
)
