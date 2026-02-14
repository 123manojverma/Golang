package main

import (
	"fmt"

	"github.com/learning/golang/auth"
	"github.com/learning/golang/user"
	"github.com/fatih/color"
)

// go mod tidy -> this command used to remove unused library 

func main() {
	auth.LoginWithCredentials("golanggyan","secret")

	session:=auth.GetSession()
	fmt.Println("Session",session)	

	user:=user.User{
		Email: "user@gmail.com",
		Name: "John Doe",
	}

	// fmt.Println(user.Email,user.Name)\
	color.Green(user.Email)
}

