package main

import "fmt"

type Operation interface {
	Add() int
	Subtract() int
}

type Number struct {
	num1   int
	num2   int
	result int
}

func (n *Number) Add() int {
	n.result = n.num1 + n.num2
	return n.result
}

func (n *Number) Subtract() int {
	n.result = n.num1 - n.num2
	return n.result
}

func NewNumber(a, b int) Operation {
	return &Number{
		num1: a,
		num2: b,
	}
}

func main() {
	n := NewNumber(10, 20)
	fmt.Printf("add: %v", n.Add())

	n1 := Number{10, 20, 0}

	a := Operation(&n1)
	fmt.Println(a.Add())

	//a is the variable of the type Operation and hold the reference to the Number struct `n1`

	/*
		a:=Operation(n1)

		This will throw error--->   Number does not implement Operation (Add method has pointer receiver)

	*/

	//if we are passing the struct to the interface, we can directly pass variable if any of the function of struct does not have the pointer reciever
	//if they have pointer reciever, then we need to pass the address of the struct varible

	//safe case:-> always pass the address of the struct variable to the interface object.

}
