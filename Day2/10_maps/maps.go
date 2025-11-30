package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main(){
	// creating map

	// m:=make(map[string]string)

	// setting an element 
	// m["name"]="golang"
	// m["area"]="backend"

	// get an element
	// fmt.Println(m["name"])

	// IMP: if key does not exists in the map then it returns zero value


	// m:=make(map[string]int)
	// m["age"]=30
	// m["price"]=300
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m,"price")
	// fmt.Println(m)

	// m:=map[string]int{"price":40,"phones":3}
	// fmt.Println(m)

	// k,ok:=m["price"]
	// fmt.Println(k)
	// if ok{
	// 	fmt.Println("all ok")
	// }else{
	// 	fmt.Println("not ok")
	// }


	m1:=map[string]int{"price":40,"phones":4}
	m2:=map[string]int{"price":40,"phones":4}

	fmt.Println(maps.Equal(m1,m2))
}