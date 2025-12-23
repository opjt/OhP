package handler

import (
	"net/http"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/user"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/log"
	"ohp/internal/pkg/token"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	log      *log.Logger
	frontUrl string
	service  *user.UserService
}

func NewUserHandler(log *log.Logger, env config.Env, service *user.UserService) *UserHandler {
	return &UserHandler{
		log:      log,
		frontUrl: env.FrontUrl,
		service:  service,
	}
}
func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/whoami", h.Whoami)

	return r
}

type resWhoami struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
}

func (h *UserHandler) Whoami(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaim, err := token.UserFromContext(ctx)
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.service.FindByEmail(ctx, userClaim.UserID)
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}
	h.log.Debug("...", "user", user)

	resp := resWhoami{
		UserID: user.ID,
		Email:  user.Email,
	}
	wrapper.RespondJSON(w, http.StatusOK, resp)

}
