package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//http get
	// resp, err := http.Get("https://google.com")
	// if err != nil {
	// 	panic(err)
	// }

	// baca body atau output
	// byteBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp.Header.Get("content-type"))

	// fmt.Println(string(byteBody))

	// implement http post / post form

	// type PostBody struct {
	// 	Name  string `json:"name"`
	// 	Email string `json:"email"`
	// }

	// postBody, _ := json.Marshal(PostBody{
	// 	Name:  "Afis",
	// 	Email: "Afis@example.com",
	// })

	// responseBody := bytes.NewBuffer(postBody)

	resp, err := http.PostForm("https://postman-echo.com/post", url.Values{
		"name":       {"admin"},
		"occupation": {"engineer"},
	})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// dial , ngetes koneksi apakah httpnya aktif atau tidak
	// conn, err := net.Dial("tcp", "dasklhjsahdjksahd.com:80")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Result :", status)
}
