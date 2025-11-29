package main

import "fmt"

const age=30
var name string = "golang"

func main(){
	// const name string="golang"
	// const age=30
	// fmt.Print(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port,host)
}