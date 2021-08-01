package main

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky-plugins/resty"
	"github.com/SkyAPM/go2sky/reporter"
	"log"
	"time"
)

func main() {
	// Use gRPC reporter for production
	// re, err := reporter.NewGRPCReporter("127.0.0.1:11800")
	re, err := reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer re.Close()

	tracer, err := go2sky.NewTracer("gin-server", go2sky.WithReporter(re))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}

	// create resty client
	client := resty.NewGoResty(tracer)
	// do something

	resp, err := client.R().Get("http://127.0.0.1:8888/ping")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(resp.Body()))

	time.Sleep(5 * time.Second)
}