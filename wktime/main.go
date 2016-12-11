package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	// proxy.OnRequest(goproxy.DstHostIs("www.baidu.com")).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// 	// log request body
	// 	log.Println("I am doing the func!")
	// 	body, _ := ioutil.ReadAll(req.Body)
	// 	log.Println(body)
	// 	return req, nil
	// })

	// this works for http, but does not work for https
	// proxy.OnRequest(goproxy.DstHostIs("www.baidu.com")).DoFunc(
	// 	func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// 		return r, goproxy.NewResponse(r,
	// 			goproxy.ContentTypeText, http.StatusForbidden,
	// 			"Don't waste your time!")
	// 	})

	proxy.OnRequest(goproxy.DstHostIs("www.baidu.com")).HijackConnect(func(req *http.Request, client net.Conn, ctx *goproxy.ProxyCtx) {
		log.Println("I am doing the func!")
		body, _ := ioutil.ReadAll(req.Body)
		log.Println(body)
	})
	log.Println("start proxy")
	log.Fatal(http.ListenAndServe(":9999", proxy))
}
