package main

import (
	"fmt"
	"time"
)

func main() {
	// timers represent a single event in the future.
	// You tell the timer how long to wait and it provides
	//  a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until
	// it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer1 fired.")

	// For waiting, time.sleep can also be used.
	// But timers can also be cancelled before firing
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired.")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
	time.Sleep(2 * time.Second)

}
