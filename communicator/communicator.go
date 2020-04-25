package communicator

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ServiceDetails struct {
	EndPoint string
	Protocol string
	Headers  map[string]string
}

//CommunicateWithService Method to communicate with Other Processes
func CommunicateWithService(service ServiceDetails, method string, urlString string, channel chan []byte) {

	// Creating New Channel for Handling Goroutine
	ch := make(chan []byte)
	switch method {

	case http.MethodGet:
		go makeGetRequest(service, urlString, ch)
	case http.MethodDelete:
		go makeDeleteRequest(service, urlString, ch)
	case http.MethodPost:
		go makePostRequest(service, urlString, ch)
	case http.MethodPatch:
		go makePatchRequest(service, urlString, ch)
	case http.MethodPut:
		go makePutRequest(service, urlString, ch)
	default:
		log.Warn("No Specified Method for Request")
	}
	channel <- (<-ch)
}

func makeGetRequest(service ServiceDetails, urlString string, ch chan []byte) {

	// Creating New Client Using Resty
	client := resty.New()

	res, err := client.R().Get(fmt.Sprintf("%s%s", service.EndPoint, urlString))

	if err != nil {
		fmt.Println("Error Occiured", err.Error())
		log.Error(err.Error())
		return
	}

	// We have the Response Data We can do whatever we want to
	fmt.Println("Response:", res.String())
	ch <- res.Body()
}

func makePostRequest(service ServiceDetails, urlString string, ch chan []byte) {
	// Creating New Client Using Resty
	client := resty.New()

	res, err := client.R().Get(fmt.Sprintf("%s%s", service.EndPoint, urlString))

	if err != nil {
		fmt.Println("Error Occiured", err.Error())
		log.Error(err.Error())
		return
	}

	// We have the Response Data We can do whatever we want to
	fmt.Println("Response:", res.String())
	ch <- res.Body()
}

func makePutRequest(service ServiceDetails, urlString string, ch chan []byte) {

	// Creating New Client Using Resty
	client := resty.New()

	res, err := client.R().Get(fmt.Sprintf("%s%s", service.EndPoint, urlString))

	if err != nil {
		fmt.Println("Error Occiured", err.Error())
		log.Error(err.Error())
		return
	}

	// We have the Response Data We can do whatever we want to
	fmt.Println("Response:", res.String())
	ch <- res.Body()
}
func makePatchRequest(service ServiceDetails, urlString string, ch chan []byte) {

	// Creating New Client Using Resty
	client := resty.New()

	res, err := client.R().Get(fmt.Sprintf("%s%s", service.EndPoint, urlString))

	if err != nil {
		fmt.Println("Error Occiured", err.Error())
		log.Error(err.Error())
		return
	}

	// We have the Response Data We can do whatever we want to
	fmt.Println("Response:", res.String())
	ch <- res.Body()
}
func makeDeleteRequest(service ServiceDetails, urlString string, ch chan []byte) {

	// Creating New Client Using Resty
	client := resty.New()

	res, err := client.R().Get(fmt.Sprintf("%s%s", service.EndPoint, urlString))

	if err != nil {
		fmt.Println("Error Occiured", err.Error())
		log.Error(err.Error())
		return
	}

	// We have the Response Data We can do whatever we want to
	fmt.Println("Response:", res.String())
	ch <- res.Body()
}
