package main

import (
	"fmt"
	"go_sugared/config"
	"go_sugared/models"

	"go_sugared/routers"
	"net/http"
)

func main() {
	fmt.Println("go sugared running.............")
	config.InitConfig()
	models.InitMongo()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Setting.Server.Port),
		Handler: router,
	}
	data1 := models.WxUser{
		Name:     "zhanan",
		Password: "12356",
	}
	// 插入数据
	models.InsertUser(&data1)

	s.ListenAndServe()
}
