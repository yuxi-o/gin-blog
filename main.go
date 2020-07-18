package main

import (
	"fmt"
	"net/http"
	"log"

	_ "github.com/gin-gonic/gin"
	"blog/pkg/setting"
	"blog/routers"
)

func main(){
	router := routers.InitRouter()

	s := &http.Server{
		Addr:			fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:		router,
		ReadTimeout:	setting.ReadTimeout,
		WriteTimeout:	setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server starting error: %v", err)
	}
}
