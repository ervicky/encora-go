package main

import (
	"log"
	"time"
)

var (
	counterLimit = 5          // limit when we want to stop go routine.
	sc           = NewCount() // initiate the global counter.
)

func main() {
	for v := 0; v < 10; v++ {
		go func(v int) {
			doublev := callDouble(v)
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}

	time.Sleep(time.Second * 10)
}
