package main

import (
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

//kita bisa menggunakan attack yang meningkat seperti ini untuk stress/capacity test

//melakukan attack selama 4 detik
//setiap detik frequency akan naik sesuai increaseValue
//nilai yang digunakan
/*
duration := 4 * time.Second
frequency := 1
increaseValue := 2
*/

/*
      / 400
     -
    /
  -
 /
- 100
*/

// load testing
// stress test
// splike test
// capacity
// concurrency
// scalability
// dst

func increaseTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	var (
		frequency     = 1
		increaseValue = 2
	)

	// setup
	// rate
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	// duration
	duration := 4 * time.Second
	// targetter
	targetter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target,
		Body:   []byte(""),
	})

	metrics = vegetaAttackIncreaseBySecond(targetter, rate, duration, increaseValue)

	return metrics
}

//frequency request akan naik sebesar increaseValue
//diawal tidak perlu menambahkan increaseValue
//contoh
//pertama 10 request
//kedua 15 request
//ketiga 20 request
//total 45 request

func vegetaAttackIncreaseBySecond(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration, increaseValue int) *vegeta.Metrics {
	var metrics vegeta.Metrics
	// TODO: answer here
	attacker := vegeta.NewAttacker()
	for i := 0; i < int(duration.Seconds()); i++ {
		if i != 0 {
			rate.Freq += increaseValue
		}

		for res := range attacker.Attack(targeter, rate, time.Second, "Increase Attack Test") {
			metrics.Add(res)
		}
	}

	metrics.Close()
	return &metrics
}
