package main

import (
	"log"
	"time"
)

func main() {
	// use channel for the throttling
	ch := make(chan bool, 5) // size 5
	
	for v := 0; v < 10; v++ {
		ch <- true // fill
		go func(v int) {
			doublev := callDouble(v)
			<-ch  // read
			log.Printf("Thread %d returned: %d", v, doublev)
		}(v)
	}
	
	// wait for go routines to finish
	for i := 0; i < cap(ch); i++ {
   		 ch <- true
	}

	time.Sleep(time.Second * 10)
}

func callDouble(v int) int {
	// adjust code to call double only up to 5 times concurrently
	return double(v)
}

// sleep for simulating any heavy background processing ?
func double(v int) int {
	time.Sleep(time.Second)
	return v * 2
}

