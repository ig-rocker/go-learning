package main

import (
	"fmt"
	"time"
)

func Multiply(n int, ch chan int) {
	fmt.Printf("value of n  %v", n)
	res := n * 23
	ch <- res
	ch <- 3434
	fmt.Println("End")

}
func main() {
	ch := make(chan int)
	go Multiply(33, ch)
	go Multiply(1, ch)
	ch1 := <-ch
	ch2 := <-ch
	ch3 := <-ch
	ch4 := <-ch
	fmt.Println("ch1", ch1)
	fmt.Println("ch2", ch2)
	fmt.Println("ch3", ch3)
	fmt.Println("ch4", ch4)

	time.Sleep(time.Second * 5)

}
