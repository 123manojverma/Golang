package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int){
	for num:=range numChan{
		fmt.Println("Processing number",num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int,num1 int ,num2 int){
	numResult:=num1+num2
	result<-numResult
}

// goroutine synchronizer
func task(done chan bool){
	defer func() {done<-true}()
	fmt.Println("processing...")
}

// <-chan it set that onlu we can receive the data can't add on it and chan<- in this we can send data doesn't receive
func emailSender(emailChan <-chan string,done chan<- bool){
	defer func(){done<-true}()
	
	for email:=range emailChan{
		fmt.Println("Sending email to",email)
		time.Sleep(time.Second)
	}
}

func main() {
	// messagechan := make(chan string)
	// messagechan <- "ping"  //blocking
	// msg := <-messagechan
	// fmt.Println(msg)

	// numChan:=make(chan int)
	// go processNum(numChan)
	// for{
	// 	numChan<-rand.Intn(100)
	// }
	// time.Sleep(time.Second*2)

	// result:=make(chan int)
	// go sum(result,4,5)
	// res:=<-result //blocking
	// fmt.Println(res)

	// done:=make(chan bool)
	// go task(done)
	// <-done // block

	// emailChan:=make(chan string,100)
	// done:=make(chan bool)
	// go emailSender(emailChan,done)
	// for i:=0;i<5;i++{
	// 	emailChan<-fmt.Sprintf("%d@gmail.com",i)
	// }
	// fmt.Println("done sending.")
	// // this is important to stop deadlock
	// close(emailChan)
	// <-done


	// emailChan<-"1@example.com"
	// emailChan<-"2@example.com"
	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)


	chan1:=make(chan int)
	chan2:=make(chan string)

	go func(){
		chan1<-10
	}()

	go func(){
		chan2<-"pong"
	}()

	for i:=0;i<2;i++{
		select{
		case chan1Val:=<-chan1:
			fmt.Println("Received data from chan1",chan1Val)
		case chan2Val:=<-chan2:
			fmt.Println("Received data from chan2",chan2Val)
		}
	}
}

