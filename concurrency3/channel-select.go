// select which works like a switch but for channels
// This program prints “from 1” every 2 seconds and “from 2” every 3 seconds.
// select picks the first channel that is ready and receives from it (or sends to it).
// If more than one of the channels are ready then it randomly picks which one to receive from.
// If none of the channels are ready, the statement blocks until one becomes available.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "form 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			ch2 <- "form 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		// time.After creates a channel and after the given duration will send the current time on it.
		// (we weren't interested in the time so we didn't store it in a variable)
		for {
			select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
