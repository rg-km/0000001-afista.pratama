package main

import (
	"fmt"
	"sync"
	"time"
)

//mutex digunakan untuk menjaga akses ke critical section
//dengan mutex kita bisa menghindari race condition
//kenapa tidak menggunakan channel ? karena kadang ada keadaan dimana menggunakan shared variabel
//lebih mudah daripada menggunakan channel
//kali ini kita menggunakan contoh sederhana
func main() {
	mu := &sync.Mutex{} // unlock mutex
	counter := 0

	//mu.Unlock() // error
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()         //terjadi blocking saat memanggil lock, jika sudah ada goroutine lain yang memanggil lock
			defer mu.Unlock() //membuka lock agar goroutine lain bisa lanjut memanggil lock
			counter += 1      // critical content
		}()

		go func() {
			mu.Lock()
			counter += 2
			mu.Unlock()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(counter)
}
