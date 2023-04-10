package main

import (
	"fmt"
	"sync"
)


func main(){
	fmt.Println("Race condition - GO")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("one R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("four R")
		m.Lock()
		score = append(score, 4)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
		fmt.Println(score)
}