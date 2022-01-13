package main

// Integrate with Grafana

import (
	"fmt"
	"net/http"
	"time"
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	fmt.Println(makeTimestamp())
	json := `[
		{"time": 1602656680028, "value": 10},
		{"time": 1602656681028, "value": 20},
		{"time": 1602656682028, "value": 30}
	]`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, json)
	})

	http.ListenAndServe(":8080", nil)
}
