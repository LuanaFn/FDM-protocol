package main

import (
	"github.com/LuanaFn/FDM-protocol/pkg/log"
	"github.com/LuanaFn/FDM-protocol/pkg/order"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Warning.Println("Error trying to load environment variables from .env file:", err)
	}

	order.HandleRequests(os.Getenv("ORDERS_ENDPOINT"))

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
