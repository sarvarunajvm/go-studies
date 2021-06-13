package main

import "fmt"

//Once the channel is closed we cannot sent the values to it.
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

	// Will chck the channel until it recieves 
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

	// Sends message to the channel 
	// 6 jobs sent to channel
    for j := 1; j <= 6; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
	// This will close the channel once its done.
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}