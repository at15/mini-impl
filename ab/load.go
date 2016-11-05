package ab

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// Load defines the workload
type Load struct {
	// The base request, attributes are set from command line arguments, like http method
	// for every new request, make a clone of it
	BaseRequest *http.Request
	// N is the total number of requests
	N int
	// C is concurrent number
	C int
}

func (l *Load) work(num int) {
	client := http.Client{}
	// TODO: QPS
	for i := 0; i < num; i++ {
		_, err := client.Do(l.BaseRequest)
		if err != nil {
			log.Print(err)
		}
	}
}

// Run block and finish all requests
func (l *Load) Run() {
	log.Print("check config")
	if l.C > l.N {
		log.Fatalf("total number %d is smaller than concurrent number %d", l.N, l.C)
	}
	requestPerWorker := l.N / l.C
	log.Printf("total request %d", l.N)
	log.Printf("request per worker %d", requestPerWorker)
	log.Printf("worker number %d", l.C)

	log.Print("issue a single request to check if server is running")
	// make a single request
	client := http.Client{}
	res, err := client.Do(l.BaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()

	log.Print("start generating load")
	var wg sync.WaitGroup
	wg.Add(l.C)
	for i := 0; i < l.C; i++ {
		go func() {
			l.work(requestPerWorker)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Print("load finished")

}
