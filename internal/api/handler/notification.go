package handler

import (
	"net/http"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/notifications"
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/log"
	"ohp/internal/pkg/token"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

// type resNoti struct {
// 	EndpointName string    `json:"endpoint_name"`
// 	Body         string    `json:"body"`
// 	IsRead       bool      `json:"is_read"`
// 	CreatedAt    time.Time `json:"created_at"`
// }

// func (h *NotiHandler) GetList(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	userClaim, err := token.UserFromContext(ctx)
// 	if err != nil {
// 		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	notis, err := h.service.GetList(ctx, userClaim.UserID)
// 	if err != nil {
// 		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	resp := make([]resNoti, len(notis))
// 	for i, noti := range notis {
// 		resp[i] = resNoti{
// 			EndpointName: noti.EndpointInfo.Name,
// 			Body:         noti.Noti.Body,
// 			IsRead:       noti.Noti.IsRead,
// 			CreatedAt:    noti.Noti.CreatedAt,
// 		}
// 	}

// 	wrapper.RespondJSON(w, http.StatusOK, resp)
// }

// 개별 알림 응답 DTO
type resNoti struct {
	ID           uuid.UUID `json:"id"` // 클라이언트가 커서로 쓸 ID
	EndpointName string    `json:"endpoint_name"`
	Body         string    `json:"body"`
	IsRead       bool      `json:"is_read"`
	CreatedAt    time.Time `json:"created_at"`
}

// 무한 스크롤 전용 응답 컨테이너
type resNotiList struct {
	Items      []resNoti  `json:"items"`
	NextCursor *uuid.UUID `json:"next_cursor"` // 다음 요청 시 사용할 ID
	HasMore    bool       `json:"has_more"`
}

func (h *NotiHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userClaim, err := token.UserFromContext(ctx)
	if err != nil {
		wrapper.RespondJSON(w, http.StatusUnauthorized, err)
		return
	}

	// 1. 쿼리 파라미터 파싱 (cursor & limit)
	cursorStr := r.URL.Query().Get("cursor")
	limitStr := r.URL.Query().Get("limit")

	var lastID *uuid.UUID
	if cursorStr != "" {
		if parsed, err := uuid.Parse(cursorStr); err == nil {
			lastID = &parsed
		}
	}

	limit := 20 // 기본값
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = min(l, 100) // 최대치 제한
	}

	notis, err := h.service.GetListWithCursor(ctx, userClaim.UserID, lastID, int32(limit))
	if err != nil {
		wrapper.RespondJSON(w, http.StatusInternalServerError, err)
		return
	}

	// 3. DTO 매핑
	items := make([]resNoti, len(notis))
	for i, noti := range notis {
		items[i] = resNoti{
			ID:           noti.Noti.ID,
			EndpointName: noti.EndpointInfo.Name,
			Body:         noti.Noti.Body,
			IsRead:       noti.Noti.IsRead,
			CreatedAt:    noti.Noti.CreatedAt,
		}
	}

	// 4. 다음 커서 결정
	var nextCursor *uuid.UUID
	hasMore := false
	if len(items) > 0 && len(items) == limit {
		// 마지막 아이템의 ID가 다음 요청의 커서가 됩니다.
		nextCursor = &items[len(items)-1].ID
		hasMore = true
	}

	wrapper.RespondJSON(w, http.StatusOK, resNotiList{
		Items:      items,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	})
}
