package main

import (
	"net/http"
	"url-sorter/internal/config"
	"url-sorter/internal/logger"
	"url-sorter/internal/router"
	"url-sorter/internal/storage"
)

func main() {
	config.MustSetEnv()
	cnf := config.MustLoad()
	log := logger.SetupLogger(cnf.ENV)
	log.Info("starting url shorter", cnf.ENV)
	stg, err := storage.New(cnf.StorageCnf)
	if err != nil {
		log.Debug(err.Error())
	}

	rt := router.NewRouter(log, stg)

	srv := &http.Server{
		Addr:         cnf.Address,
		Handler:      rt,
		ReadTimeout:  cnf.HTTPServer.Timeout,
		WriteTimeout: cnf.HTTPServer.Timeout,
		IdleTimeout:  cnf.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
