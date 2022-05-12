package main

import (
	"fmt"
	"sync"
	"time"
)

// implement singleton
// memastikan bahwa instance dari class singleton hanya ada 1

type db struct {
	conn string
}

var (
	instance *db
	mu       sync.Mutex
)

func ConnectDB() *db {
	// thread safe
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		fmt.Println("create new instance")
		instance = &db{
			conn: "localhost:3306",
		}
	}

	fmt.Println("return instance")
	return instance
}

var once sync.Once

func ConnectDBSafe() *db {
	once.Do(func() {
		instance = &db{
			conn: "localhost:3306",
		}
	})

	return instance
}

func main() {
	// goroutine 100
	//var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i, "=", ConnectDB())
		}(i)
	}

	time.Sleep(2 * time.Second)
}
