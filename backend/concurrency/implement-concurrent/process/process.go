package process

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func StartProcess(records [][]string) {
	// syncronous
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for _, record := range records {
		GetResponseWebsite(record, &client)
	}

	// goroutine
}

// crawler
// goroutine
// channel
// set max worker / di limit goroutine
func GetResponseWebsite(record []string, client *http.Client) {
	// http

	url := "https://" + record[1]

	log.Println("request GET", url)
	req, err := http.NewRequest(http.MethodGet, "https://"+record[1], nil)
	if err != nil {
		log.Fatal("error :", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	log.Println("response GET", url)

	body, _ := ioutil.ReadAll(resp.Body)

	f, err := os.Create("response-web/" + record[1] + ".html")
	if err != nil {
		log.Fatal("error :", err)
	}

	_, err = f.Write(body)
	if err != nil {
		log.Fatal("error :", err)
	}
}
