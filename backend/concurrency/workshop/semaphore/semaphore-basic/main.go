package main

import (
	"fmt"
	"log"
	"time"
)

type Semaphore interface {
	Acquire(i int)
	Release(i int)
}

type semaphore struct {
	semC chan struct{}
}

func newSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire(i int) {
	log.Println("send token", i)
	s.semC <- struct{}{}
}

func (s *semaphore) Release(i int) {
	log.Println("send back token", i)
	<-s.semC
}

//kita bisa lihat hanya ada 10 goroutine yang berjalan dalam satu waktu
func main() {
	//

	semaphore := newSemaphore(10) // kita ingin hanya ada 10 akses dalam satu waktu

	//semaphore2 := semaphore.NewWeighted()

	for i := 1; i <= 30; i++ {
		semaphore.Acquire(i)
		go func(i int) {
			defer semaphore.Release(i)
			longRunningProcess(i)
		}(i)
	}
	//kapan terjadi blocking pada program ini ?
	// TODO: answer here
	// karena menggunakan buffer channel, maka akan blocking ketika cap udah terisi full, dan baru lanjut ketika channelnya di receive

	// time.Sleep
	// waitGroup
	// channel / channel buffer
}

func longRunningProcess(taskID int) {
	fmt.Println("blocking")
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(1 * time.Second) // melakukan sesuatu
	fmt.Println("after blocking")
}
