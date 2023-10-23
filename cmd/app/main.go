package main

import (
	"log"
	auth2 "rchir7/internal/auth"
	"rchir7/internal/config"
	redis2 "rchir7/internal/db"
	"rchir7/internal/http"
	"rchir7/internal/http/handlers"
)

func main() {
	c := config.NewConfig()
	rdb := redis2.NewRedis(c)
	auth := auth2.NewTokenHandler(c.Secret)
	handler := handlers.NewHandler(auth, rdb)
	router := http.NewRouter(handler)

	log.Fatal(router.R.Run("0.0.0.0:" + c.ServerPort))
}
