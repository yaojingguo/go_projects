package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)


func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func httpGet(url string) {
	req, _ := http.NewRequest("GET", url, nil)

	var start time.Time
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		TLSHandshakeStart: func() {
			start = time.Now()
			log.Print("TLS handshake started")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			end := time.Now()
			log.Printf("elapsed: %v", end.Sub(start))
			// log.Printf("state: %#v, err: %s", state, err)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	httpGet("https://httpbin.org/get")
	httpGet("https://httpbin.org")
}
