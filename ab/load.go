package ab

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Load struct {
	// The base request, attributes are set from command line arguments, like http method
	// for every new request, make a clone of it
	BaseRequest *http.Request
}

func (l *Load) Run() {
	log.Print("start generating load")
	// make a single request
	client := http.Client{}
	res, err := client.Do(l.BaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
