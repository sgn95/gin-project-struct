package main

import (
	"demo/routers"
	"demo/setting"
	"fmt"
	"net/http"
)

func init() {
	setting.SetUp()
}

func main() {
	ginHandler := routers.InitRouter()
	endPort, _ := setting.ToInt("HttpPort")
	endPoint := fmt.Sprintf(":%d", endPort)
	server := &http.Server{
		Addr:    endPoint,
		Handler: ginHandler,
	}
	server.ListenAndServe()
}
