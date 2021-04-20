package order

import (
	"fmt"
	"github.com/LuanaFn/FDM-protocol/pkg/log"
	"io/ioutil"
	"net/http"
)

var businessHost string

func create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(businessHost, "application/json", nil)
	if err != nil {
		log.Error.Printf("error connecting to order service at %s for consumer %s: %+v", businessHost, r.Host, err)
		handleError("error connecting to order service", w)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error.Printf("error sending response from %s to consumer %s: %+v", businessHost, r.Host, err)
		handleError("error retrieving response from provider", w)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Warning.Printf("Error closing body: %+v", err)
	}

	_, err = w.Write(bodyBytes)
	if err != nil {
		log.Error.Printf("error writing response from %s to consumer %s: %+v", businessHost, r.Host, err)
		handleError("error retrieving response from provider", w)
	}
	log.Info.Printf("Order created: %s", string(bodyBytes))
}

func handleError(msg string, w http.ResponseWriter) {
	errorResponse := fmt.Sprintf("{\"error\":\"%s\"}", msg)
	w.WriteHeader(http.StatusInternalServerError)
	_, wErr := w.Write([]byte(errorResponse))
	if wErr != nil {
		log.Error.Panicf("error communicating error to consumer: \nError:%+v\nOriginal error:%s", wErr, msg)
	}
}

/**
HandleRequests register the handlers for the APIs in this package
To expose the APIs you must run http.ListenAndServe after calling this
*/
func HandleRequests(host string) {
	businessHost = host

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			create(w, r)
		default:
			log.Error.Print("error: invalid request ", r.Method)
			handleError("invalid request", w)
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
