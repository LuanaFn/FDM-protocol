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
	c := make(chan int)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Error.Fatal(err)
		}

		c <- 1
	}()

	log.Debug.Println("Listening on port 8080")
	log.Debug.Print(<-c)
}
