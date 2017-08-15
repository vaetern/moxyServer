package servingStrategies

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

type ServeAndStoreStrategy struct{

}

func NewServeAndStoreStrategy() ServeAndStoreStrategy{
	return ServeAndStoreStrategy{}
}

func (s ServeAndStoreStrategy) Start(operationPort *string, verbose *bool){

	proxyInstance := goproxy.NewProxyHttpServer()
	proxyInstance.Verbose = *verbose
	log.Fatal(http.ListenAndServe(":"+*operationPort, proxyInstance))

	proxyInstance.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})
}
