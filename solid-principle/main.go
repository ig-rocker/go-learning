package main

import "fmt"

//1. Single Responsibilty Principle

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	//Log to the file
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	//log to the console
}

/* in the above two struct, they have their functionality to log
tomorrow if i have to make some change in the file logger concept, i know exactly where I have to look for!
*/

//=====================================//

/*2. Open Close Principle
In Go, this one is acheivable with the help of the interface and composition.
*/
type PaymentGateway interface {
	Pay(amount float32) bool
}

type CreditCardGateway struct{}

func (g *CreditCardGateway) Pay(amount float32) bool {
	//concept
	return true
}

type WalletGateway struct{}

func (w *WalletGateway) Pay(amount float32) bool {
	//concept
	return false
}

/* So in the above, we can able to create a object Creditcard and can pass to the Payment inerface as it is open for the extension
not allowed to modify the interface. */

//===========================================//

/* 3. Liskov Substitution Principle
In this object of the superclass is replacable of the object of the subclass wothout affecting the correctness of the program
In Go, this is achievable using interface */

func IsPaid(p PaymentGateway) {
	fmt.Println("the amount is Paid: ", p.Pay(32))
}

//=================================================================//

/* 4. Interface Segregation Principle
It state that client shold not be forced to depend on the interface they do not use.
In Go, this means creating the smaller,focused interface */

type Read interface{}

type Write interface{}

//If we write it in the same interface, it will not follow the ISA principle

//======================================================

/* 5. Dependency Inversion Principle
In this the high level module does not need to depend on the low level module.
Instead both should depends on the  abstraction */

type Keybaord interface{}
type Monitor interface{}

type Computer struct{}

func main() {

}
