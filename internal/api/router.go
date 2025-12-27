package api

import (
	"ohp/internal/api/handler"
	middle "ohp/internal/api/middleware"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/token"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

func NewRouter(
	subscriptionHandler *handler.SubscriptionHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	endpointHandler *handler.EndpointHandler,
	apiHandler *handler.ApiHandler,

	tokenProvider *token.TokenProvider,
	env config.Env,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middle.CorsMiddleware(env.FrontUrl))

	r.Mount("/api", apiHandler.Routes())
	r.Mount("/auth", authHandler.Routes())

	r.Group(func(r chi.Router) {
		r.Use(middle.AuthMiddleware(tokenProvider))
		r.Mount("/subscriptions", subscriptionHandler.Routes())
		r.Mount("/users", userHandler.Routes())
		r.Mount("/endpoints", endpointHandler.Routes())
	})

	return r
}

var routeModule = fx.Module("router",
	fx.Provide(
		handler.NewSubscriptionHandler,
		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewEndpointHandler,

		// API
		handler.NewApiHandler,
	),

	fx.Provide(NewRouter),
)
