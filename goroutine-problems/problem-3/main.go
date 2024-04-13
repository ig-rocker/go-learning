package main

import (
	"fmt"
	"sync"
)

//problem statement
/*
You have given two array and you need to create a goroutine function
that gonna return a new array having sum of the two array.

*/

// func sumOfArray(arr1 [][]int32, arr2 [][]int32, wg *sync.WaitGroup) {
// 	var arr3 [3][4]int32

// 	for i, val1 := range arr1 {
// 		for j, _ := range val1 {
// 			arr3[i][j] = arr1[i][j] + arr2[i][j]
// 		}
// 	}

// 	fmt.Printf("Sum of the array is %+v", arr3)
// 	wg.Done()
// }



func sumOfArray2(arr1 [][]int32, arr2 [][]int32, wg *sync.WaitGroup){

	var wg2 sync.WaitGroup
	var arr3[3][4]int32

	wg.Add(len(arr1))

	for i, val:=range arr1{	
		go func(){
			for j,_:=range val{
				arr3[i][j]=arr1[i][j]+arr2[i][j]
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
		fmt.Printf("Sum of the array is %+v", arr3)

	wg.Done()

}

func main() {

	arr := [][]int32{{1, 3, 4, 5}, {3, 5, 2, 5}, {4, 1, 9, 3}}

	arr2 := [][]int32{{1, 3, 4, 5}, {3, 5, 2, 5}, {4, 1, 9, 3}}

	var wg sync.WaitGroup

	wg.Add(1)
	// sumOfArray(arr, arr2, &wg)
	sumOfArray2(arr, arr2, &wg)
	wg.Wait()
}
