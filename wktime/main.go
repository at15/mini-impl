package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		// log request body
		log.Println("I am doing the func!")
		body, _ := ioutil.ReadAll(req.Body)
		log.Println(body)
		return req, nil
	})
	log.Println("start proxy")
	log.Fatal(http.ListenAndServe(":9999", proxy))
}
