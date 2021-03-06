package ab

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
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
	// Q is the QPS of every worker
	Q int
	// T is the timeout
	T int
	// TODO: why use channel instead of array
	results chan *result
	report  *report
}

type result struct {
	err      error
	code     int
	duration time.Duration
}

func (l *Load) work(num int) {
	var t <-chan time.Time
	// FIXED: this solve the connect: cannot assign requested address problem
	tr := &http.Transport{}
	client := http.Client{Transport: tr, Timeout: time.Duration(l.T) * time.Second}

	// FIXME: this does not reuse connection
	// client := http.Client{Timeout: time.Duration(l.T) * time.Second}
	if l.Q > 0 {
		log.Printf("QPS is %d", l.Q)
		t = time.Tick(time.Duration(1e6/l.Q) * time.Microsecond)
	}
	for i := 0; i < num; i++ {
		// TODO: do we need to put initial value into t? I think not
		if l.Q > 0 {
			<-t
		}
		r := new(http.Request)
		*r = *l.BaseRequest

		s := time.Now()

		// res, err := client.Do(l.BaseRequest)
		res, err := client.Do(r)
		// TODO: it seems forgetting this will also cause file descriptor problem
		// res.Body.Close()
		t := time.Now()
		code := 0
		if err != nil {
			log.Print(err)
		} else {
			io.Copy(ioutil.Discard, res.Body)
			// FIXME: now the request is canceled
			res.Body.Close()
		}
		if res != nil {
			code = res.StatusCode
		}
		result := &result{duration: t.Sub(s), err: err, code: code}
		l.results <- result
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
	client := http.Client{Timeout: time.Duration(l.T) * time.Second}
	res, err := client.Do(l.BaseRequest)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()

	l.results = make(chan *result, l.N)
	defer close(l.results)

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

	l.report = newReport(l.results)
	l.report.finalize()
}

func (l *Load) PrintReport() {
	l.report.print()
}
