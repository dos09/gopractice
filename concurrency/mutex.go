package concurrency

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) inc() {
	c.mutex.Lock()
	c.count += 1
	c.mutex.Unlock()
}

func (c *Counter) dec() {
	c.mutex.Lock()
	c.count -= 1
	c.mutex.Unlock()
}

func (c *Counter) readValue() int{
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func incCounter(c *Counter, ch chan int) {
	for i := 0; i < 1000; i++ {
		c.inc()
		if i % 100 == 0 {
			fmt.Println(c.readValue())
		}
	}
	ch <- 1 //end indication for the main thread
}

func decCounter(c *Counter, ch chan int) {
	for i := 0; i < 1000; i++ {
		c.dec()
		if i % 100 == 0 {
			fmt.Println(c.readValue())
		}
	}
	ch <- 1 //end indication for the main thread
}

func RunMutex() {
	fmt.Println(" * RunMutex")
	counter := Counter{}
	ch := make(chan int, 2)
	go incCounter(&counter, ch)
	go decCounter(&counter, ch)

	for i := 0; i < 2; i++ {
		<-ch //wait for the two threads to finish
	}
	fmt.Printf("counter value = %d\n", counter.count)
}
