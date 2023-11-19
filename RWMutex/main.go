package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mux := &sync.RWMutex{}

	go writeLoop(m, mux)
	go readLoop(m, mux)
	go readLoop(m, mux)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 5; i++ {
			mux.Lock()
			fmt.Println("write lock ", i)
			m[i] = i
			mux.Unlock()
			fmt.Println("write unlock ", i)

		}
	}
}

func readLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		fmt.Println("read lock ")

		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.RUnlock()
		fmt.Println("read unlocking ")

	}
}
