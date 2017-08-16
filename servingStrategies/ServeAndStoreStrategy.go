package servingStrategies

import (
	"github.com/elazarl/goproxy"
	commStoreService "github.com/vaetern/moxyServer/communicationBodyService"
	"github.com/vaetern/moxyServer/servingStrategies/ComLog"
	"net/http"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"bytes"
	"io/ioutil"
)

const headerRequestKey = "Moxy-Request-Key"
const headerSoapAction = "Soapaction"

type ServeAndStoreStrategy struct {
}

func NewServeAndStoreStrategy() (strat ServeAndStoreStrategy) {
	return strat
}

func (s ServeAndStoreStrategy) Start(operationPort *string, verbose *bool) {
	log.Println("Proxy server start")
	comLogCh := make(chan ComLog.CommunicationLog)

	proxyInstance := goproxy.NewProxyHttpServer()

	proxyInstance.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxyInstance.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		rqBodyBytes, _ := ioutil.ReadAll(r.Body)
		rqBodyString := string(rqBodyBytes)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(rqBodyBytes))
		comHashedBody := commStoreService.NewComHashedBody(rqBodyString)
		r.Header.Set(headerRequestKey, comHashedBody.Output)
		return r, nil
	})

	proxyInstance.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			commLog := ComLog.CommunicationLog{}

			commLog.Target = ctx.Req.Header.Get(headerSoapAction)

			commLog.ResponseKey = string(ctx.Req.Header.Get(headerRequestKey))

			rsBodyBytes, _ := ioutil.ReadAll(r.Body)
			rsBodyString := string(rsBodyBytes)
			commLog.ResponseBody = string(rsBodyString)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(rsBodyBytes))

			comLogCh <- commLog
			log.Println("ch<-")

			return r
		})
	proxyInstance.Verbose = *verbose

	commBody := commStoreService.NewStoreService()
	commBody.ProcessStoring(comLogCh)

	log.Fatal(http.ListenAndServe(":" + *operationPort, proxyInstance))
}
