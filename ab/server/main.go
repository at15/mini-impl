package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// simply print the http protocol
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Protocol: %s", r.Proto)
}

func main() {
	log.Print("start server")
	port := "8000"
	// NOTE: will have permission denied
	// port := "800"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// handle os signal
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	go func() {
		sig := <-s
		log.Printf("signal %v", sig)
		log.Print("Press Ctrl + C again to exit")
		<-s
		// NOTE: without this, singal interrupt will only show once and ctrl + c will never work ...
		os.Exit(1)
	}()

	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	// mux.Handler()
	// server.Handler(mux)
	server.Handler = mux
	server.Addr = "localhost:" + port
	// err := server.ListenAndServe("localhost:"+port, nil)
	err := server.ListenAndServe()

	// https://gist.github.com/denji/12b3a568f092ab951456
	// openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
	// ln -sf server.rsa.key server.key
	// ln -sf server.rsa.crt server.crt
	// err := http.ListenAndServeTLS("localhost:"+port, "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
