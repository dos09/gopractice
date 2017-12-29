package concurrency

import (
	"fmt"
	"time"
)

func RunDefaultSelect() {
	fmt.Println(" * RunDefaultSelect")
	practiceWithoutDefault()
	practiceWithDefault()
}

func practiceWithoutDefault() {
	fmt.Println("practiceWithoutDefault")
	c := make(chan int, 10)
	go sendOneNumDelayed(c)
	select {
		//will block here on empty channel
		case i := <- c:
			fmt.Printf("received %d\n", i)
	}
}

func practiceWithDefault() {
	fmt.Println("practiceWithDefault")
	c := make(chan int, 10)
	go sendOneNumDelayed(c)
	select {
		case i := <- c:
			fmt.Printf("received %d\n", i)
		//instead of blocking on the above case when the channel is empty
		//executes the default case
		default:
			fmt.Printf("default\n")
	}
}

func sendOneNumDelayed(c chan int) {
	time.Sleep(1 * time.Second)
	c <- 100
	fmt.Println("sent number")
}