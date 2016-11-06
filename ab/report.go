package ab

import (
	"fmt"
	"time"
)

type report struct {
	fastest float64
	slowest float64

	results chan *result
	total   time.Duration
}

func newReport(results chan *result) *report {
	return &report{
		fastest: 99999.99,
		slowest: 0.00,
		results: results,
	}
}

func (r *report) finalize() {
	// FIXME: this will stuck forever?
	// for res := range r.results {
	// 	if r.fastest > res.duration.Seconds() {
	// 		r.fastest = res.duration.Seconds()
	// 	}
	// 	if r.slowest < res.duration.Seconds() {
	// 		r.slowest = res.duration.Seconds()
	// 	}
	// }
	for {
		select {
		case res := <-r.results:
			if r.fastest > res.duration.Seconds() {
				r.fastest = res.duration.Seconds()
			}
			if r.slowest < res.duration.Seconds() {
				r.slowest = res.duration.Seconds()
			}
		default:
			return
		}
	}
}

func (r *report) print() {
	fmt.Printf("slowest %4.4fs\n", r.slowest)
	fmt.Printf("fastest %4.4fs\n", r.fastest)
}
