package main

import "fmt"

// enumerated types

// type OrderStatus int
type OrderStatus string

const (
	// Received OrderStatus=iota
	// Confirmed
	// Prepared
	// Delivered

	// for string
	Received OrderStatus="received"
	Confirmed OrderStatus = "confirmed"
	Prepared OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to",status)
}

func main(){
	// changeOrderStatus(Delivered)
	changeOrderStatus(Received)
}