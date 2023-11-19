package main

import (
	"fmt"
	"sync"
	"time"
)

type game struct {
	mutex sync.RWMutex
	score int
}

func (g *game) GetScore() {
	g.mutex.RLock()

	for i := 0; i < 4; i++ {
		fmt.Println("Điểm số:", g.score)
	}

	g.mutex.RUnlock()

}

func (g *game) IncreaseScore() {
	// Khóa khóa rmutex trước khi tăng biến 'score'
	g.mutex.Lock()
	fmt.Println("write lock")
	// Tăng biến 'score'
	g.score++

	// Mở khóa khóa rmutex sau khi tăng biến 'score'
	g.mutex.Unlock()
	fmt.Println("write unlocking")

}

func main() {
	g := game{
		score: 0,
	}

	go func() {
		for i := 0; i < 10; i++ {
			g.IncreaseScore()
			time.Sleep(1 * time.Second)
		}
	}()

	// Tạo các goroutines để đọc điểm số
	for i := 0; i < 10; i++ {
		go g.GetScore()
		time.Sleep(2 * time.Second)

	}

	// Chờ các goroutines hoàn thành
	time.Sleep(6 * time.Second)
}
