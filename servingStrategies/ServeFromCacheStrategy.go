package servingStrategies

import (
	"github.com/elazarl/goproxy"
	"net/http"
	comService "github.com/vaetern/moxyServer/communicationBodyService"
	"io/ioutil"
	"bytes"
	"log"
)

const contentType = "text/xml"

type ServeFromCacheStrategy struct {
}

func NewServeFromCacheStrategy() ServeFromCacheStrategy {
	return ServeFromCacheStrategy{}
}

func (s ServeFromCacheStrategy) Start(operationPort *string, verbose *bool) {
	log.Println("Mocker server start")
	cachedService := comService.NewComCachedService()

	proxyInstance := goproxy.NewProxyHttpServer()

	proxyInstance.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxyInstance.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

			rqBodyBytes, _ := ioutil.ReadAll(r.Body)
			rqBodyString := string(rqBodyBytes)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(rqBodyBytes))
			comHashedBody := comService.NewComHashedBody(rqBodyString)

			target := r.Header.Get(HeaderSoapAction)

			responseBody, err := cachedService.GetCachedBodyFor(comHashedBody, target)

			log.Println(target)

			if err != nil {
				log.Println(err)
			} else {
				log.Println("Y - served from cache")
			}

			return r, goproxy.NewResponse(r, contentType, http.StatusOK, responseBody)
		})

	proxyInstance.Verbose = *verbose

	log.Println("-> ready to serve")

	log.Fatal(http.ListenAndServe(":" + *operationPort, proxyInstance))
}
