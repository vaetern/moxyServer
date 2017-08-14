package main

import (
	"runtime"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

const MaxProcesses = 8

const ServingPort = "19501"

func main() {

	runtime.GOMAXPROCS(MaxProcesses)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":"+ServingPort, proxy))

	proxy.OnRequest().DoFunc(
		func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			r.Header.Set("X-GoProxy","yxorPoG-X")
			return r,nil
		})
}
