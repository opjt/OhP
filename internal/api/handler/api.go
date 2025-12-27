package handler

import (
	"context"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/push"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/log"

	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	log     *log.Logger
	service *push.PushService
}

func NewApiHandler(
	log *log.Logger,
	env config.Env,

	service *push.PushService,
) *ApiHandler {
	return &ApiHandler{
		log:     log,
		service: service,
	}
}
func (h *ApiHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/push/{token}", wrapper.WrapJson(h.Push, h.log.Error, wrapper.RespondJSON))

	return r
}

func (h *ApiHandler) Push(ctx context.Context, _ interface{}) (interface{}, error) {
	token := chi.URLParamFromCtx(ctx, "token")
	h.log.Info("...", "token", token)

	if err := h.service.Push(ctx, token); err != nil {
		return nil, err
	}

	return nil, nil
}
