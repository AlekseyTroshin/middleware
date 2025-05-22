package load

import (
	storageMap "hexlet-auth/internal/storage/storage-map"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger, storage *storageMap.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const OP = "handlers.url.save.New"

		log := log.With(
			slog.String("op", OP),
		)
	}
}
