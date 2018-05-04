package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {

	path := ctx.Path()
	fmt.Println(string(path))
	ipaddr := ctx.RemoteIP()

	ctx.Write([]byte("Hello World"))
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q %s", ctx.RequestURI(), ipaddr)

}

func main() {

	if err := fasthttp.ListenAndServe(":8000", requestHandler); err != nil {
		log.Fatalf("Error in server %s", err)
	}
}
