package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		ch2 <- 23
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			fmt.Println("entered into channel 1")

		case <-ch2:
			fmt.Println("entered into channel 2")
		}
	}

}
