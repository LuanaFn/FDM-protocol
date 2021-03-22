package main

import (
	"github.com/LuanaFn/FDM-protocol/configs"
	"github.com/LuanaFn/FDM-protocol/pkg/log"
	"github.com/LuanaFn/FDM-protocol/pkg/order"
	"net/http"
)

func main() {
	err := configs.Config.Load("")
	if err != nil {
		log.Error.Panic(err)
	}
	log.Info.Printf("Running with config: %+v", configs.Config)
	order.HandleRequests()
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Error.Panic("error opening order service: ", err)
	}
}
