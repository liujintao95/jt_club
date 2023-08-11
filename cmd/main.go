package main

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/routes"
	"JT_CLUB/pkg/cache"
	"JT_CLUB/pkg/db"
	"JT_CLUB/pkg/log"
)

func main() {
	conf.InitConfig()
	log.InitLogger()
	db.InitDataBase()
	cache.InitCache()
	r := routes.NewRouter()
	if err := r.Run(); err != nil {
		log.Logger.Fatal(err.Error())
	}
}
