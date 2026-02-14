package auth

import "fmt"

// If any func start with small letter it only accessible in same directory 
// If you want to access that func in other directory the func first letter must be capital 
func LoginWithCredentials(username string,password string){
	fmt.Println("login user using",username,password)
}