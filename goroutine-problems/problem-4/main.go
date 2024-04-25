package main

import ("fmt"
	"time"
	)

// var ch chan int

// func Producer(wg *sync.WaitGroup){
// 	for i:=0; i<10; i++{
// 		ch<-i
// 		time.Sleep(time.Second)
// 	}
// 	wg.Done()
// }


// func Consumer(ch <-chan int){
// 	for {
// 		fmt.Println(<-ch)
// 	}
// }


// func main(){
// 	wg:=sync.WaitGroup{}
// 	wg.Add(1)
// 	ch=make(chan int)
// 	go Producer(&wg)
// 	go Consumer(ch)
// 	wg.Wait()
// }














var (ch1 chan int
	ch2 chan string)

func main(){
	ch1=make(chan int)
	ch2=make(chan string)
	go func(){
		ch1<-1
		// ch2<-"sdsfc"
		close(ch2)
	}()
	

	time.Sleep(time.Second)
	
	fmt.Println("ch: ",<-ch1)
	val,ok:=<-ch2
	if ok{
		fmt.Println(val)
	}

	}