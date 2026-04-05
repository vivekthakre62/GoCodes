package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	price     float32
	status    string
	items     []string
	createdAt string
}

// constructor function for order struct
func newOrder(id string, price float32, status string) {
	myOrder3 := order{
		id:     id,
		price:  price,
		status: status,
	}

	fmt.Println(myOrder3)
}

func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}
func (o order) displayItems() []string {
	return o.items
}

func main() {

	//inline struct only one time use

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)

	myOrder := order{
		id:        "cs-101",
		price:     100.0,
		status:    "Panding",
		items:     []string{"Macbook Pro", "iPhone 18", "ipad pro"},
		createdAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Println(myOrder)

	myOrder2 := order{}
	myOrder2.id = "CS-102"
	myOrder2.changeStatus("Pending")
	fmt.Println(myOrder2)
}
