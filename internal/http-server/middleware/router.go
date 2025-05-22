package middleware

import (
	"encoding/json"
	"hexlet-auth/internal/lib/api/cookies"
	storageMap "hexlet-auth/internal/storage/storage-map"
	"log/slog"
	"net/http"
)

type RequestData struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

func SignIn(
	log *slog.Logger,
	cookies *cookies.Session,
	storage *storageMap.Storage,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const OP = "internal.http-server.middleware.SignIn"

		log := log.With(
			slog.String("op", OP),
		)

		if r.Method != http.MethodPost {
			log.Info("Method not POST")
			return
		}

		session := cookies.GetCookieHandler(log, "session_id", r)

		if session == "" {
			log.Info("Session - not found")
		} else {
			log.Info("Session - we have session")
			next(w, r)
			return
		}

		var requestData RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			log.Error("failed decoder to body", "error", err)
			return
		}

		if requestData.User == "" || requestData.Email == "" {
			log.Error("need to enter right data user/email", "error", err)
			return
		}

		_, err = storage.Load(requestData.Email)
		if err != nil {
			log.Error("need to authorization", "error", err)
			return
		}

		cookies.SetCookieHandler(log, "session_id", w)

		next(w, r)
	}
}

func SignUp(
	log *slog.Logger,
	cookies *cookies.Session,
	storage *storageMap.Storage,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const OP = "internal.http-server.middleware.SignUp"

		log := log.With(
			slog.String("op", OP),
		)

		if r.Method != http.MethodPost {
			log.Info("method not POST")
			return
		}

		next(w, r)
	}
}
