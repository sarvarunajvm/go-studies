package main

import (
	"fmt"
	"time"
)

//Tickers are for when you want to do something
// repeatedly at regular intervals.
func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Started ticker every 500 millis
	time.Sleep(1600 * time.Millisecond)

	// Tickers can be stopped like timers.
	// Once a ticker is stopped it won’t receive any
	// more values on its channel. We’ll stop ours after 1600ms.
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
