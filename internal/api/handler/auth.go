package handler

import (
	"net/http"
	"ohp/internal/domain/auth"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/log"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	log      *log.Logger
	frontUrl string
	service  *auth.AuthService
}

func NewAuthHandler(log *log.Logger, env config.Env, service *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		log:      log,
		frontUrl: env.FrontUrl,
		service:  service,
	}
}
func (h *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/github/callback", h.OauthGithubCallback)

	return r
}

func (h *AuthHandler) OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Redirect(w, r, h.frontUrl+"/login", http.StatusFound)
		return
	}

	//  GitHub API로 사용자 정보 가져오기
	userProfile, err := h.service.OauthGithubFlow(code)
	if err != nil {
		http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}
	h.log.Info("Github user profile", "user", userProfile)

	// // DB 작업 (Upsert: 있으면 가져오고 없으면 생성)
	// user, err := h.userService.FindOrCreate(userProfile)
	// if err != nil {
	// 	http.Error(w, "Database error", http.StatusInternalServerError)
	// 	return
	// }

	// // JWT 생성
	// token, err := h.service.GenerateJWT(user) // 유저 ID 등을 Payload에 담음
	// if err != nil {
	// 	http.Error(w, "JWT Generation failed", http.StatusInternalServerError)
	// 	return
	// }

	// 프론트엔드로 JWT 전달 (Cookie 또는 Query Parameter)
	// 보안상 HttpOnly Cookie를 사용.
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "testoken",
		Path:     "/",
		HttpOnly: true,  // 자바스크립트 접근 방지
		Secure:   false, // HTTPS 권장
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24, // 1일
	})

	http.Redirect(w, r, h.frontUrl+"/", http.StatusFound)
}
