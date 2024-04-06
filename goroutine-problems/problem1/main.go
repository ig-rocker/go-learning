package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Problem: We have to create 4 goroutines which have
1-> it need to generate random number
2 -> find out the number is even/odd
3 -> print even number
4 -> print odd number
*/

func GenerateRandomNum() {
	num := rand.Intn(100)
	fmt.Println("num", num)

	ch := make(chan bool)
	go IsEvenOdd(num, ch)
	<-ch
}

func IsEvenOdd(n int, ch chan bool) {
	ch1 := make(chan int)
	if n%2 == 0 {
		go EvenNumber(n, ch1)
		<-ch
	} else {
		go OddNumber(n, ch1)
		<-ch
	}
	ch <- true
}

func EvenNumber(n int, ch chan int) {
	fmt.Println("Number is Even", n)
	ch <- 1
}
func OddNumber(n int, ch chan int) {
	fmt.Println("Number is Odd", n)
	ch <- 1
}

func main() {
	go GenerateRandomNum()
	time.Sleep(time.Second * 2)
}
