package schedules

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/your-user/scheduling-api-prototype-go/internal/httpx"
)

func ListHandler(store *Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpx.OK(r.Context(), w, map[string]any{"schedules": store.List()})
	})
}

type createReq struct {
	Title string `json:"title"`
	Time  string `json:"time"` // RFC3339 string
}

func CreateHandler(store *Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req createReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.Error(r.Context(), w, http.StatusBadRequest, "bad_request", "invalid JSON body")
			return
		}
		if req.Title == "" || req.Time == "" {
			httpx.Error(r.Context(), w, http.StatusBadRequest, "bad_request", "title and time are required")
			return
		}
		t, err := time.Parse(time.RFC3339, req.Time)
		if err != nil {
			httpx.Error(r.Context(), w, http.StatusBadRequest, "bad_request", "time must be RFC3339")
			return
		}
		created := store.Create(req.Title, t)
		httpx.Created(r.Context(), w, map[string]any{"schedule": created})
	})
}
