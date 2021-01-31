package http_client_heimdall

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/gojektech/heimdall/v6/hystrix"
)

func TestRequestGet(t *testing.T) {
	// Create a new HTTP client with a default timeout
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get("http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestRequestGetWithDo(t *testing.T) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Create an http.Request instance
	req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestRequestGetWithHystrix(t *testing.T) {
	fallbackFn := func(err error) error {
		fmt.Println("error1:", err)
		return err
	}

	// Create a new hystrix-wrapped HTTP client with the command name, along with other required options
	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Millisecond),
		hystrix.WithCommandName("google_get_request"),
		hystrix.WithHystrixTimeout(1000*time.Millisecond),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
		hystrix.WithFallbackFunc(fallbackFn),
	)

	// Use the clients GET method to create and execute the request
	res, err := client.Get("https://www.google.com.hk/?gws_rd=ssl", nil)
	if err != nil {
		fmt.Println("error2:", err)
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
