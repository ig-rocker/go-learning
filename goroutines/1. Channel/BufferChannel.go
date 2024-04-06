package main

import "fmt"

/*
	In Buffered channel, we have one more parameter which tells the size of the buffer.
	Initially, the sender send the data and wait until the receiver receive it and channel can only save one value at a time.
	Now, sender can send the data till the buffered size and then wait till the Reciever starts consuming the data.


*/

func MultiplyTriple(arr []int, ch chan int) {
	for _, ele := range arr {
		ch <- ele * 3
	}
}

func BufferedChannel() {

	fmt.Println("Started consuming the buffered goroutine")
	arr := []int{3, 5, 2, 7, 9}
	ch1 := make(chan int, len(arr))

	go MultiplyTriple(arr, ch1)

	for i := 0; i < len(arr); i++ {
		fmt.Println("value of multiple", <-ch1)
	}

}
