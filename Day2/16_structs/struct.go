package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// struct embedding
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// var order order OR

	// if you don't set any field, default value is zero value
	// int->0, float->0, string->"", bool->false
	// myOrder:=order{
	// 	id:"1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.createdAt=time.Now()
	// fmt.Println(myOrder.status)
	// fmt.Println("Order struct",myOrder)

	// myOrder2:=order{
	// 	id:"2",
	// 	amount: 100,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder2.changeStatus("confirmed")
	// fmt.Println(myOrder2.getAmount())
	// fmt.Println(myOrder2)

	// myOrder:=newOrder("1",30.50,"received")
	// fmt.Println(myOrder.amount)

	// if want to use struct only one time than declare it inside the main function
	// language := struct{
	// 	name string
	// 	isGood bool
	// }{"golang",true}

	// fmt.Println(language)

	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		// customer: newCustomer,
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	newOrder.customer.name="robin"
	fmt.Println(newOrder)

}
