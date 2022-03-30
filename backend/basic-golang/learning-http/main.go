package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// dial , ngetes koneksi apakah httpnya aktif atau tidak
	conn, err := net.Dial("tcp", "dasklhjsahdjksahd.com:80")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :", status)
}
