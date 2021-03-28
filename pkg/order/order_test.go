package order

import (
	"bytes"
	"fmt"
	"github.com/LuanaFn/FDM-protocol/configs"
	"github.com/LuanaFn/FDM-protocol/pkg/log"
	"github.com/cucumber/godog"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var orderServerMock *httptest.Server
var createOrderHttpResp *http.Response
var orderReceived = false
var callbackUrl string
var externalVendorProducts bool
var businessMsg = []byte("{\"msg\":\"order was received\"}")

func before(_ *godog.Scenario) {
	orderServerMock = httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug.Println("Order was received in business mock service")
			orderReceived = true
			_, err := w.Write(businessMsg)
			if err != nil {
				log.Error.Panic("Error writing answer from mock server: ", err)
			}
		}),
	)
	log.Debug.Println("Orders mocked business server up")
	configs.Config.Business.Endpoints = map[string]string{
		"order": orderServerMock.URL,
	}
	configs.Config.ServiceHost = "http://localhost:8080"
	log.Debug.Println("Orders mocked business configured")

	HandleRequests()
	go func() {
		err := http.ListenAndServe("localhost:8080", nil)
		if err != nil {
			log.Error.Println("error opening order service: ", err)
		}
	}()
	log.Debug.Println("Orders service up")
}

func after(_ *godog.Scenario, _ error) {
	orderServerMock.Close()
}

func iDoNotAddProductsFromExternalVendors() error {
	externalVendorProducts = false
	return nil
}

func iDoNotProvideACallbackURL() error {
	callbackUrl = ""
	return nil
}

func iOrderAListOfProductsFromCurrentVendor() error {
	client := &http.Client{}
	host := configs.Config.ServiceHost + "/orders"
	req, err := http.NewRequest("POST", host, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	createOrderHttpResp, err = client.Do(req)
	if err != nil {
		return err
	}
	log.Debug.Printf("order sent to %s", host)
	return nil
}

func iReceiveConfirmationThatMyOrderWasSent() error {
	if createOrderHttpResp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(createOrderHttpResp.Body)
		if err != nil {
			return fmt.Errorf("error reading response: %+v", err)
		}
		err = createOrderHttpResp.Body.Close()
		if err != nil {
			log.Warning.Printf("Error closing body: %+v", err)
		}
		if bytes.Equal(bodyBytes, businessMsg) {
			return nil
		} else {
			return fmt.Errorf("business message was not received.\nMessage received: %b\nMessage expected: %b", bodyBytes, businessMsg)
		}
	}
	return fmt.Errorf("error: create order service responded with non-200 status: %s", createOrderHttpResp.Status)
}

func theOrderIsSentToTheOrderBusinessService() error {
	if orderReceived {
		return nil
	} else {
		return fmt.Errorf("error: the order was not received by the order business service at \"%s\"", configs.Config.Business.Endpoints["order"])
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(before)
	ctx.AfterScenario(after)
	ctx.Step(`^I do not add products from external vendors$`, iDoNotAddProductsFromExternalVendors)
	ctx.Step(`^I do not provide a callback URL$`, iDoNotProvideACallbackURL)
	ctx.Step(`^I order a list of products from current vendor$`, iOrderAListOfProductsFromCurrentVendor)
	ctx.Step(`^I receive confirmation that my order was sent$`, iReceiveConfirmationThatMyOrderWasSent)
	ctx.Step(`^the order is sent to the order business service$`, theOrderIsSentToTheOrderBusinessService)
}
