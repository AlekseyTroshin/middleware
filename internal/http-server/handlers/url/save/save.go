package save

import (
	"fmt"
	"hexlet-auth/internal/lib/api/render"
	storageMap "hexlet-auth/internal/storage/storage-map"
	"log/slog"
	"net/http"
)

type Request struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

func New(log *slog.Logger, storage *storageMap.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const OP = "handlers.url.save.New"

		log := log.With(
			slog.String("op", OP),
		)

		var req Request

		err := render.DecoderJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", "error", err)
			render.JSON(w, r)
			return
		}

		user, err := storage.Save(req.User, req.Email)
		if err != nil {
			log.Error("have user", "error", err)
			return
		}

		log.Info(fmt.Sprintf("signup new user: %s %s", user.Name, user.Email))
	}
}
