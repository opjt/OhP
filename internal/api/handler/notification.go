package handler

import (
	"net/http"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/notifications"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/log"
	"ohp/internal/pkg/token"
	"time"

	"github.com/go-chi/chi/v5"
)

type NotiHandler struct {
	log     *log.Logger
	service *notifications.NotiService
}

func NewNotiHandler(
	log *log.Logger,
	env config.Env,

	service *notifications.NotiService,
) *NotiHandler {
	return &NotiHandler{
		log:     log,
		service: service,
	}
}
func (h *NotiHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetList)

	return r
}

type resNoti struct {
	EndpointName string    `json:"endpoint_name"`
	Body         string    `json:"body"`
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
}

func (h *NotiHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaim, err := token.UserFromContext(ctx)
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}
	notis, err := h.service.GetList(ctx, userClaim.UserID)
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]resNoti, len(notis))
	for i, noti := range notis {
		resp[i] = resNoti{
			EndpointName: noti.EndpointInfo.Name,
			Body:         noti.Noti.Body,
			IsRead:       noti.Noti.IsRead,
			CreatedAt:    noti.Noti.CreatedAt,
		}
	}

	wrapper.RespondJSON(w, http.StatusOK, resp)
}
