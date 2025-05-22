package cookies

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	sessions map[string]string
}

func New() *Session {
	return &Session{
		sessions: map[string]string{},
	}
}

func (s *Session) SetCookieHandler(log *slog.Logger, name string, w http.ResponseWriter) {
	const OP = "internal.lib.api.cookies.SetCookieHandler"

	log = log.With(
		slog.String("op", OP),
	)

	val := uuid.New().String()
	cookie := &http.Cookie{
		Name:    name,
		Value:   val,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
	s.sessions[val] = "TRUE"
	log.Info("cookies are set")
}

func (s *Session) GetCookieHandler(log *slog.Logger, name string, r *http.Request) string {
	const OP = "internal.lib.api.cookies.GetCookieHandler"
	log = log.With(
		slog.String("op", OP),
	)

	cookie, err := r.Cookie(name)
	if err != nil {
		log.Error("request cookie not found", "status", http.StatusNotFound)
		return ""
	}

	val, ok := s.sessions[cookie.Value]
	if !ok {
		log.Error("sessions cookie not found", "status", http.StatusNotFound)
		return ""
	}

	return val
}

func (s *Session) DeleteCookieHandler(log *slog.Logger, name string, w http.ResponseWriter) {
	const OP = "internal.lib.api.cookies.DeleteCookieHandler"
	log = log.With(
		slog.String("op", OP),
	)

	getCookie, ok := s.sessions[name]
	if !ok {
		log.Error("there are no cookies to delete ", http.StatusNotFound)
		return
	}

	cookie := &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	delete(s.sessions, getCookie)
	log.Info("cookie deletion successful")

}
