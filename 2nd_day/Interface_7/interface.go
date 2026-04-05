package main

import "fmt"

// interface

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	getway paymenter
}
type razorpay struct {
}

//if i want to integrate other api then
type other struct{}

func (p payment) makePayment() {
	// razorpayPayment := razorpay{}
	//never use like that because method alway should be work on open/close principle
	// otherPayment := other{}
	// otherPayment.pay(12000)
	p.getway.pay(12000)
}

func (r razorpay) pay(amount float32) {
	fmt.Println("amount paid by razorpay ", amount)
}
func (o other) pay(amount float32) {
	fmt.Println("amount paid by other api", amount)
}

func main() {
	razorpayGet := razorpay{}
	paymentcreate := payment{
		getway: razorpayGet,
	}
	paymentcreate.makePayment()
}
