package process

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func AsyncStartProcess(Worker int, records [][]string, fileOutput string) {
	// thread pool
	var wg sync.WaitGroup
	var ch = make(chan []string)

	// initialize
	for i := 0; i < Worker; i++ {
		wg.Add(1)
		go GOGetRespCodeWebsite(&wg, ch, fileOutput)
	}

	for _, record := range records {
		ch <- record
	}

	close(ch)
	wg.Wait()

	log.Println("Async process done")
}

// thread pool

// loop sampai 100
// 100
// 1<- 200 <- 1 1 1 1 1 1 1  1 1
// 1<- 199 <- 1 1 1 1 1 1 1 1  1 1
// 1, 1 <- 99 <- 1 1 1 1  1 1 1  1 1

// set pool

func GOGetRespCodeWebsite(wg *sync.WaitGroup, ch chan []string, fileOutput string) {
	defer wg.Done()

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for c := range ch {
		url := "https://" + c[1]

		log.Println("request GET", url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal("error :", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		log.Println("response GET", url)

		f, err := os.OpenFile(fileOutput, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatal("error :", err)
			os.Create(fileOutput)
		}

		csvWriter := csv.NewWriter(f)

		respCode := strconv.Itoa(resp.StatusCode)

		err = csvWriter.Write([]string{c[0], c[1], respCode})
		if err != nil {
			log.Fatal("error :", err)
		}

		csvWriter.Flush()
		f.Close()

		continue
	}

}

func GoGetResponseWebsite(wg *sync.WaitGroup, ch chan []string) {
	defer wg.Done()

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for c := range ch {
		url := "https://" + c[1]

		log.Println("request GET", url)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal("error :", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		log.Println("response GET", url)

		body, _ := ioutil.ReadAll(resp.Body)

		f, err := os.Create("response-web/" + c[1] + ".html")
		if err != nil {
			log.Fatal("error :", err)
		}

		_, err = f.Write(body)
		if err != nil {
			log.Fatal("error :", err)
		}

		f.Close()
		resp.Body.Close()

		continue
	}
}
