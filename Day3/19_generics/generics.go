package main

import "fmt"

// We can use any or interface{} if we don't know the incoming type
// func printSlice[T int|string|bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable, V string](items []T,name V) {
	for _, item := range items {
		fmt.Println(item,name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T any] struct{
	elements []T
}

func main(){
	// nums:=[]int{1,2,3,4}
	// names:=[]string{"golang","typescript"}
	vals:=[]bool{true,false,true}
	printSlice(vals,"john")
	// printStringSlice(names)

	myStack:=stack[string]{
		// elements: []int{1,2,3},
		elements: []string{"golang","java"},
	}
	fmt.Println(myStack)
}