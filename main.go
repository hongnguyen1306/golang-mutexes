package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type account struct {
	mutex  sync.Mutex
	amount float64 // Số dư tài khoản
}

func (acc *account) Deposit(sum float64) {
	acc.mutex.Lock() // Khóa khóa mutex trước khi sửa đổi 'amount'
	time.Sleep(time.Duration(rand.Int31n(250)) * time.Millisecond)
	acc.amount += sum  // Cập nhật số dư tài khoản
	acc.mutex.Unlock() // Mở khóa khóa mutex sau khi sửa đổi 'amount'
}

func (acc *account) Withdraw(sum float64) {
	acc.mutex.Lock() // Khóa khóa mutex trước khi sửa đổi 'amount'
	time.Sleep(time.Duration(rand.Int31n(250)) * time.Millisecond)
	acc.amount -= sum  // Cập nhật số dư tài khoản
	acc.mutex.Unlock() // Mở khóa khóa mutex sau khi sửa đổi 'amount'
}

func main() {
	// Tạo tài khoản
	acc := account{
		amount: 1000,
	}

	go func() {
		acc.Deposit(100)
	}()

	go func() {
		acc.Withdraw(50)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(acc.amount)
}
