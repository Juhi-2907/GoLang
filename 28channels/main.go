package main

import (
	"fmt"
	"sync"
)

// channels are a way that multiple go routine can talk to each other

func main(){
	fmt.Println("Channels in - GO")
	
	// myCh := make(chan int)

	myCh := make(chan int, 2) //buffered channel

	// // I will only allow you to pass a value only when someone is listening to you
	// myCh <- 5
	// fmt.Println(<-myCh)

	// fmt.Println(<-myCh) //won't work because no value is there
	// myCh <- 5

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup){ //receiving from channel
		val,isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh,wg)

	go func(ch chan<- int, wg *sync.WaitGroup){ //sender in channel
		// close(ch)
		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(myCh,wg)

	wg.Wait()
}
