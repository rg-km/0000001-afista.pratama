package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	// rate
	rate := vegeta.Rate{
		Freq: 100,
		Per:  time.Second,
	}
	// duration
	duration := 5 * time.Second

	// targetter / target
	targetter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8000/",
		Body:   []byte(""),
	})

	// setup attacker
	attacker := vegeta.NewAttacker()

	// set metric
	metric := &vegeta.Metrics{}

	// run attacker
	for res := range attacker.Attack(targetter, rate, duration, "load test vegeta basic") {
		fmt.Printf("%s, res code : %d\n", string(res.Body), res.Code)
		metric.Add(res)
	}

	metric.Close()

	// print result
	report := vegeta.Report(metric)

	fmt.Println(report)
}
