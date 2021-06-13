package main

import (
	"fmt"
	"time"
)


// Timers are for when you want to do something 
//once in the future
func main() {

	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it
	// sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the
	// timer before it fires. Here’s an example of that.
	stop2 := timer2.Stop()
	if stop2 {
		// The first timer will fire ~2s after we start the program,
		// but the second should be stopped before
		// it has a chance to fire.
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
