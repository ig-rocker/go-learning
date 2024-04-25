package main

import "fmt"

func AddNum(a, b int) int {
	return a + b
}

func main() {
	a := AddNum(4, 5)
	fmt.Println("sum= ",a )
}