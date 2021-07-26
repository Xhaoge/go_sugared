package main

import (
	"fmt"
	"go_sugared/config"
	"go_sugared/pkg/logging"
	"go_sugared/routers"
	"go_sugared/routers/api"
	"net/http"
)

func main() {
	logging.Info("go sugared running.............")
	router := routers.InitRouter()
	if err := api.InitDatabase(); err != nil {
		logging.Error("init database error: %s", err)
	}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ConfigGetServicePort()),
		Handler: router,
	}

	s.ListenAndServe()
}
