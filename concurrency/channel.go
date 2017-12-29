package concurrency

import (
	"fmt"
	"time"
)

func RunChannel() {
	fmt.Println(" * RunChannel")
	practiceChannel()
	practiceBufferedChannel()
	practiceChannelClose()
}

func practiceChannel() {
	fmt.Println("practiceChannel started")
	s := []int{1, 6, 4, 3, 2, 0}
	sHalf := len(s) / 2
	c := make(chan int)
	go sumItems(s[:sHalf], c)
	go sumItems(s[sHalf:], c)

	x, y := <-c, <-c //will block untill both go routines write in the channel
	fmt.Println(x, y, x+y)
	fmt.Println("practiceChannel finished")
}

func sumItems(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to channel c
}

func practiceBufferedChannel() {
	fmt.Println("practiceBufferedChannel started")
	itemsCount := 10
	c := make(chan int, 5)
	go writeToChannel(c, itemsCount)
	var num int
	for i := 0; i < itemsCount; i++ {
		num = <- c
		fmt.Printf("read %d\n", num)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("practiceBufferedChannel finished")
}

func writeToChannel(c chan int, itemsCount int) {
	for i := 0; i < itemsCount; i++ {
		c <- i
		fmt.Printf("wrote %d\n", i)
	} 
	fmt.Println("writeToChannel finished")
}

func practiceChannelClose() {
	fmt.Println("practiceChannelClose started")
	c := make(chan int, 10)
	go sendSomeNumbers(c)
	//one way
//	for i := range c {
//		fmt.Printf("read1 %d\n", i)
//	}
	//another way
	for {
		num, ok := <- c
		if !ok {
			break
		}
		fmt.Printf("read2 %d\n", num)
	}
	
	fmt.Println("practiceChannelClose finished")
}

func sendSomeNumbers(c chan int) {
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
}
