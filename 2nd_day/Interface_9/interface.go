package main

import "fmt"

type paymenter interface {
	payment(amount float64)
}

type payment struct {
	getway paymenter
}

func (p payment) createPayment() {
	p.getway.payment(12000)
}

type upi struct {
}

func (u upi) payment(amount float64) {
	fmt.Println("Payment by using upi ", amount)
}

type creditCard struct {
}

func (cr creditCard) payment(amount float64) {
	fmt.Println("Payment by using credit card ", amount)
}

type cash struct {
}

func (c cash) payment(amount float64) {
	fmt.Println("Payment by using credit card ", amount)
}
func main() {
	getWays := cash{}
	Paygetway := payment{
		getway: getWays,
	}
	Paygetway.createPayment()
}
