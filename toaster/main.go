package main

import (
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

var count int64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Don't panic! For the " + strconv.Itoa(int(atomic.AddInt64(&count, 1))) + "th time"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
