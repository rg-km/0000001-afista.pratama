package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// tipe data semaphore
// type Semaphore interface {
// 	Acquire()
// 	Release()
// }

// type semaphore struct {
// 	semC chan struct{} // zero memory
// }

// // encapsulation / factory
// func NewSemaphore(maxProcc int) Semaphore {
// 	return &semaphore{
// 		semC: make(chan struct{}, maxProcc),
// 	}
// }

// req token
// func (s *semaphore) Acquire() {
// 	s.semC <- struct{}{}
// }

// mengembalikan token
// func (s *semaphore) Release() {
// 	<-s.semC
// }

func main() {
	// sem := NewSemaphore(5)

	// process := 20

	// wait := make(chan bool, 1)

	// for i := 1; i <= process; i++ {
	// 	sem.Acquire()
	// 	go func(i int) {
	// 		defer sem.Release()
	// 		longRunningProcess(i)

	// 		if i == process {
	// 			close(wait)
	// 		}
	// 	}(i)
	// }

	// <-wait

	// package semaphore
	var worker int64 = 5

	ctx := context.TODO()

	sem := semaphore.NewWeighted(worker) // 5

	// main
	// process 1
	// go1 - go5
	// 				go6-go10
	// 							go11-go15
	//										go16-go20
	// stop, process done

	for i := 0; i < 20; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)

			longRunningProcess(i)

		}(i)
	}

	if err := sem.Acquire(ctx, worker); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v", err)
	}

}

func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(500 * time.Millisecond) // melakukan sesuatu
}
