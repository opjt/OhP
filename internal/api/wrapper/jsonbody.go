package wrapper

import (
	"context"
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIResponse{
		Code:    status,
		Success: status >= 200 && status < 300,
	}

	if errObj, ok := payload.(error); ok {
		resp.Error = errObj.Error()
	} else {
		resp.Data = payload
	}

	_ = json.NewEncoder(w).Encode(resp)
}
func WrapJson[T any](
	handler func(context.Context, T) (interface{}, error),
	logger func(msg string, keyvals ...any),
	respond func(w http.ResponseWriter, status int, data any),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto T
		if r.Body != http.NoBody {
			if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
				logger("parse json error", "err", err)
				respond(w, http.StatusBadRequest, err)
				return
			}
		}
		res, err := handler(r.Context(), dto)
		if err != nil {
			logger("handler error", "err", err)
			respond(w, http.StatusBadRequest, err)
			return
		}
		respond(w, http.StatusOK, res)
	}
}
