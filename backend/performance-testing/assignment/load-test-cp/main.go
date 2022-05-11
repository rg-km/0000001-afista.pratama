package main

import (

	// TODO: answer here
	"encoding/json"
	"strconv"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

//Baca README untuk tau jumlah request yang perlu dilakukan dan targetnya
//untuk durasi cukup gunakan satu detik

//menambahkan movie baru
//untuk data yang dikirim adalah JSON
//gunakan struct Movie diatas, cukup gunakan field episode dan name
//ID sudah auto increment
func addMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	// setup
	// frequency
	frequency := 10
	// duration
	duration := 1 * time.Second
	// targetter

	body, _ := json.Marshal(Movie{
		Episode: 1,
		Name:    "Movie 1",
	})

	targetter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target + "/movie",
		Body:   body,
	})

	metrics = vegetaAttack(targetter, frequency, duration)

	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here

	freq := 25
	duration := 1 * time.Second

	targets := []vegeta.Target{}

	// id random 1 - 25

	for i := 1; i <= 25; i++ {

		id := strconv.Itoa(i)

		targets = append(targets, vegeta.Target{
			Method: "GET",
			URL:    target + "/movie/" + id,
			Body:   []byte(""),
		})
	}

	targeter := vegeta.NewStaticTargeter(targets...)
	metrics = vegetaAttack(targeter, freq, duration)

	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	freq := 20
	duration := time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target + "/movies",
	})

	metrics = vegetaAttack(targeter, freq, duration)

	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, frequency int, duration time.Duration) *vegeta.Metrics {
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
