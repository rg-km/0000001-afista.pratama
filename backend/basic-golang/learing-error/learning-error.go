package main

import (
	"errors"
	"fmt"
)

func main() {
	// error basic

	// 2 compile error, runtime error
	//var number int = "nama saya afis"
	// runtime error
	// fmt.Println(divide(10, 0))

	// convert string ke int
	// num, err := strconv.Atoi("asdaskj!#$#") // "0" - "9"
	// var err => type interface error
	// method : Error() => return string
	// if err != nil {
	// 	log.Panic(err)
	// }
	// handling error
	// fmt.Println("hasil", num)

	// custom error
	// var err1 = errors.New("error happen")
	// var err2 = fmt.Errorf("error in line %v, is happen, %s %t", 29, "error", true)

	// fmt.Println(err1.Error())

	// result, err := divide(10, 0)
	// // 0 , terjadi error
	// if err == ErrDivideZero {
	// 	errInformation := err.Error()
	// 	fmt.Println(errInformation)

	// panic / fatal / fmt.Println()
	// fmt.Println(err)
	// log.Panic(err)

	// errUnwrap := errors.Unwrap(err) // error yang diunwrap
	// if errUnwrap != nil {
	// 	fmt.Println(errUnwrap.Error())
	// }

	// } else {
	// 	fmt.Println("hasil", result)
	// }

	// if errors.Is(err, ErrDivideZero) { // true / false
	// 	fmt.Println("error is:", err)
	// }

	// if errors.As(err, &<memori dari si error A>) {
	// fmt.Println
	// }

	// passing by reference // memori error A
	// var errDivZero

	var errMsg = errors.New("cannot divide zero")

	// fmt.Println("hello")
	// fmt.Println("hello2")
	result, err := divide(10, 0)
	// handlePanic(err)

	if err != nil {
		if errors.As(err, &errMsg) {
			fmt.Println(err)
		}
	}

	fmt.Println(result)
	// fatal, munculkan error, terus exit terminal (nyetop terminal)

	// panic, munculkan error, hentikan eksekusi / goroutine error
	// panic => bisa di recover

	// fmt.Println("hello3")
	// fmt.Println("hello4")

}

// func handlePanic(err error) {
// 	defer func() {
// 		if panic := recover(); panic != nil {
// 			fmt.Println(panic)
// 		}
// 	}()

// 	panic(err) // stop disini
// }

// eksekusi code paling terakhir sebelum func selesai
func divide(number1, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("cannot divide zero")
	} else {
		return number1 / number2, nil
	}

	// recover
	// defer func() { // IIFE
	// 	if panic := recover(); panic != nil {
	// 		fmt.Println(panic) // error
	// 	}
	// }()

	// panic => error runtime
	// terjadi error, tampilkan error dan  eksekusi berhenti

}

// Error() => dapetin informasi error
// errors.New("error happer") => ngebuat error, dengan tipe data error

// sentinel
// error wrapping
// var ErrDivideZero = fmt.Errorf("error: %w", errors.New("cannot divide by zero"))

// c, golang
// verb => %v, %t, %s, %w => convert string

// tidak ada error handling
// func divide(number1 int, number2 int) int {
// 	return number1 / number2
// }
