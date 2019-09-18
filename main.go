package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome to Api!")
}

func Cetak(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		next(ctx)
		ctx.WriteString(fmt.Sprint(" -] Accessed from: ", string(ctx.RequestURI())))
	}
}

func main() {
	r := router.New()
	r.GET("/", Index)

	log.Fatal(fasthttp.ListenAndServe(":6600", Cetak(r.Handler)))
}
