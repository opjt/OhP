package handler

import (
	"context"
	"net/http"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/endpoint"
	"ohp/internal/pkg/log"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type EndpointHandler struct {
	log     *log.Logger
	service *endpoint.EndpointService
}

func NewEndpointHandler(log *log.Logger, service *endpoint.EndpointService) *EndpointHandler {
	return &EndpointHandler{
		log:     log,
		service: service,
	}
}
func (h *EndpointHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", wrapper.WrapJson(h.Add, h.log.Error, wrapper.RespondJSON))
	r.Get("/", h.GetList)
	r.Delete("/", wrapper.WrapJson(h.Delete, h.log.Error, wrapper.RespondJSON))
	return r
}

type reqAddEndpoint struct {
	ServiceName string `json:"serviceName"`
}

func (h *EndpointHandler) Add(ctx context.Context, req reqAddEndpoint) (interface{}, error) {

	h.log.Info("req", "sub", req)
	if err := h.service.Add(ctx, req.ServiceName); err != nil {
		return nil, err
	}

	return "success", nil
}

type resListEndpoint struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Token  string    `json:"token"`
	Active bool      `json:"active"`
}

func (h *EndpointHandler) GetList(w http.ResponseWriter, r *http.Request) {

	endpoints, err := h.service.List(r.Context())
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}

	result := make([]resListEndpoint, 0, len(endpoints))
	for _, endpoint := range endpoints {
		result = append(result, resListEndpoint{
			ID:     endpoint.ID,
			Name:   endpoint.Name,
			Token:  endpoint.Token,
			Active: true,
		})
	}

	wrapper.RespondJSON(w, http.StatusOK, result)
}

type reqDeleteEndpoint struct {
	Token string `json:"token"`
}

func (h *EndpointHandler) Delete(ctx context.Context, req reqDeleteEndpoint) (interface{}, error) {

	if err := h.service.Remove(ctx, req.Token); err != nil {
		return nil, err
	}

	return "success", nil

}
