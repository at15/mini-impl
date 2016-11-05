package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/at15/mini-impl/ab"
)

func main() {
	log.Print("start ab")
	cpus := runtime.NumCPU()
	log.Printf("cpu num %d", cpus)
	r, err := http.NewRequest("GET", "http://localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
	total := 1000
	// total := 10000

	concurent := 10
	// qps := 500
	qps := 50
	// qps := 0

	l := ab.Load{BaseRequest: r, N: total, C: concurent, Q: qps}
	l.Run()

}
