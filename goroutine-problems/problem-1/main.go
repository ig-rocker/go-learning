package main

import (
	"fmt"
	"sync"
)

var (
	even chan int
	odd  chan int
	ch   chan int
	wg   sync.WaitGroup
)

func GenerateNum() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func IsEvenOrOdd() {
	for {
		num := <-ch
		if num%2 == 0 {
			even <- num
		} else {
			odd <- num
		}
	}
}

func EvenNum() {
	for {
		num := <-even
		fmt.Println("Number is even", num)
	}
}

func OddNum() {
	for {
		fmt.Println("Number is odd", <-odd)
	}
}

func main() {

	even = make(chan int)
	odd = make(chan int)
	ch = make(chan int, 5)
	wg.Add(1)
	go GenerateNum()
	go IsEvenOrOdd()
	go EvenNum()
	go OddNum()
	// time.Sleep(time.Second * 10)
	wg.Wait()
}
