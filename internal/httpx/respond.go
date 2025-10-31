package httpx

import (
	"context"
	"encoding/json"
	"net/http"
)

type envelope struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  *apiError   `json:"error,omitempty"`
	Meta   meta        `json:"meta,omitempty"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type meta struct {
	RequestID string `json:"request_id,omitempty"`
}

func write(ctx context.Context, w http.ResponseWriter, status int, data interface{}, err *apiError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{
		Status: http.StatusText(status),
		Data:   data,
		Error:  err,
		Meta:   meta{RequestID: RequestIDFromCtx(ctx)},
	})
}

func OK(ctx context.Context, w http.ResponseWriter, data interface{}) {
	write(ctx, w, http.StatusOK, data, nil)
}

func Created(ctx context.Context, w http.ResponseWriter, data interface{}) {
	write(ctx, w, http.StatusCreated, data, nil)
}

func Error(ctx context.Context, w http.ResponseWriter, status int, code, message string) {
	write(ctx, w, status, nil, &apiError{Code: code, Message: message})
}
