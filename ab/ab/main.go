package main

import (
	"log"
	"net/http"

	"github.com/at15/mini-impl/ab"
)

func main() {
	log.Print("start ab")
	r, err := http.NewRequest("GET", "http://localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
	l := ab.Load{BaseRequest: r}
	l.Run()

}
