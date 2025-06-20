package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a timer that will expire after 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// Wait for the timer to expire
	fmt.Println("Timer 1 started:", time.Now().Format(time.Kitchen))
	<-timer1.C
	fmt.Println("Timer 1 expired:", time.Now().Format(time.Kitchen))

	// Create a timer and cancel it before it expires
	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("This should not print")
	}()

	// Cancel the timer
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped successfully")
	}

	// Use time.After for a simpler way to use timers
	fmt.Println("\nTime.After example started:", time.Now().Format(time.Kitchen))
	<-time.After(3 * time.Second)
	fmt.Println("Time.After completed:", time.Now().Format(time.Kitchen))
}
