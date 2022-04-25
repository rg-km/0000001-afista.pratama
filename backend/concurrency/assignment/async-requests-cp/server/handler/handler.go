package handler

import (
	"log"
	"net/http"
	"time"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only GET work", http.StatusMethodNotAllowed)
		return
	}
	time.Sleep(5 * time.Millisecond)

	log.Println("GET, hi")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hi"))
}
