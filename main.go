package main

import (
	"fmt"
	"go_sugared/routers"
	"net/http"
	//"github.com/gin-gonic/gin"
	//"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	s.ListenAndServe()
}
