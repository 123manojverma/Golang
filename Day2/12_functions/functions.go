package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func getLanguage() (string,string,bool){
// 	return "golang","javascript",true
// }

// func processIt(fn func(a int) int){
// 	fn(1)
// }

func processIt() func(a int) int{
	return func(a int)int {
		return 4
	}
}

func main() {
	// result := add(3, 5)
	// fmt.Println(result)
	// lang1,lang2,lang3:=getLanguage()
	// fmt.Println(lang1,lang2,lang3)
	// fmt.Println(getLanguage())

	// lang1,lang2,_:=getLanguage()
	// fmt.Println(lang1,lang2)

	// fn:=func (a int) int {
	// 	return 2
	// }
	// processIt(fn)

	fn:=processIt()
	fmt.Println(fn(6))
}