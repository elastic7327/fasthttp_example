package main_test

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestFunction(t *testing.T) {

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://www.google.com")

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	client.Do(req, resp)

	bodyBytes := resp.Body()

	fmt.Println(string(bodyBytes))
}
