package main

import (
	"context"
	"fmt"
	"url-sorter/internal/config"
	"url-sorter/internal/logger"
	"url-sorter/internal/storage"
	"url-sorter/internal/storage/database"
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
	url, err := stg.UseCase.SaveURL(context.Background(), &database.SaveURLParams{Url: "123", Alias: "1", ID: 1})
	if err != nil {
		log.Debug(err.Error())
	}
	fmt.Println(url, stg)
}
