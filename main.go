package main

import (
	"fmt"
	"go_sugared/config"
	"go_sugared/routers"
	"go_sugared/routers/api"
	"net/http"
)

func main() {
	fmt.Println("go sugared running.............")
	config.InitConfig()
	router := routers.InitRouter()
	api.InitMongo()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ConfigGetServicePort()),
		Handler: router,
	}

	s.ListenAndServe()
}
