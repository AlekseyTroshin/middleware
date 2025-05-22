package main

import (
	"fmt"
	"hexlet-auth/internal/config"
	"hexlet-auth/internal/http-server/handlers/url/load"
	"hexlet-auth/internal/http-server/handlers/url/save"
	"hexlet-auth/internal/http-server/middleware"
	"hexlet-auth/internal/lib/api/cookies"
	"hexlet-auth/internal/lib/config/logger"
	storageMap "hexlet-auth/internal/storage/storage-map"
	"net/http"
)

const Port = 1234

func main() {
	cfg := config.MustConfig()

	log := logger.SetupLogger(cfg.Env)

	session := cookies.New()

	storage := storageMap.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/sign-in", middleware.SignIn(log, session, storage, load.New(log, storage)))
	mux.HandleFunc("/sign-up", middleware.SignUp(log, session, storage, save.New(log, storage)))
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1", storage)
		fmt.Println("2", session)
	})

	port := fmt.Sprintf(":%d", Port)
	log.Info("server starting on port " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		return
	}
}

func funcSet(w http.ResponseWriter, req *http.Request) {
	_, err := w.Write([]byte("\nHello 11111\n"))
	if err != nil {
		return
	}
	return
}
