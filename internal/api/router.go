package api

import (
	"ohp/internal/api/handler"
	middle "ohp/internal/api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

func NewRouter(pushHandler *handler.PushHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middle.CorsMiddleware)

	r.Mount("/push", pushHandler.Routes())

	return r
}

var routeModule = fx.Module("router",
	fx.Provide(
		handler.NewPushHandler,
	),

	fx.Provide(NewRouter),
)
