package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type DataHTTP struct {
	Address string   `json:"address"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
	Name    string   `json:"name"`
	Tinggi  float64  `json:"tinggi"`
}

func (d DataHTTP) String() string {
	return fmt.Sprintf("alamat : %s, umur : %v", d.Address, d.Age)
}

func main() {
	url := "http://localhost:8080/json-data"

	resp, err := http.Get(url)
	handleErr(err)

	bodyData, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	// unmarshal => perlu nyiapin struct
	var data DataHTTP

	err = json.Unmarshal(bodyData, &data)
	handleErr(err)

	fmt.Println(data)
}
