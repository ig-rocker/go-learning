package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
write a producer and consumer in goroutine

Producer will generate the random number at every second
Consumer will print that number after multiplying by 2.
*/
func producer(ch chan int32) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		ch <- int32(num)
		fmt.Println("Produced number",num)
		// time.Sleep(time.Second)
	}
}

func consumer(ch chan int32) {
	for{
		fmt.Println("Consumer is ", <-ch)
	}
}

func main() {
	ch := make(chan int32,3)
	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 10)
}
