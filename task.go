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

// Start function.
// N is processor number
// X is queue length. Channel max buffer
func Start(N int, X int) {

	runtime.GOMAXPROCS(N)
	queue := make(chan data, X)

	// Imitation of data input
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
		настоящий код должен быть таким поидее:
		 x := <-ch
		*** do work with x for some time ***
	*/

	// Придется сначала ждать, и только потом читать из канала
	// иначе буфер канала никогда не заполнится
	time.Sleep(time.Second * 5)
	<-ch
	fmt.Println("COMPLETED")
}
