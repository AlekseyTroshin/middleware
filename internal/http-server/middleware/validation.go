package middleware

import (
	"log/slog"
	"net/http"
)

type RequestData struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

func Validate(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
