package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	Start(2, 10)

}

type data struct{}

// Start function
func Start(N int, X int) {

	runtime.GOMAXPROCS(N)

	queue := make(chan data, X)

	var input string

	fmt.Println("Type anything and press enter")

	for {
		fmt.Scan(&input)
		x := data{}

		select {
		case queue <- x:
			fmt.Println("Task has been received")
			go process(queue)
		default:
			fmt.Println("The queue is full, please wait")
		}

	}

}

func process(ch chan data) {
	/*
		REAL CODE MUST BE:
		 x := <-ch
		*** do work with x for some time ***
	*/

	time.Sleep(time.Second * 5)
	<-ch
	fmt.Println("COMPLETED")
}
