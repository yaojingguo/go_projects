package main

import (
	"context"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky-plugins/resty"
	"github.com/SkyAPM/go2sky/reporter"
	"log"
	"time"
)

var tracer *go2sky.Tracer

func main() {
	var err error
	var re go2sky.Reporter

	re, err = reporter.NewGRPCReporter("localhost:11800")
	//re, err := reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("new reporter error: %v", err)
	}
	defer re.Close()

	tracer, err = go2sky.NewTracer("test", go2sky.WithReporter(re))
	if err != nil {
		log.Fatalf("create tracer error: %v", err)
	}
	first(tracer)
	time.Sleep(1 * time.Second)
}

func first(tracer *go2sky.Tracer) {
	ctx := context.Background()
	span, _, _ := tracer.CreateLocalSpan(ctx, go2sky.WithOperationName("six"))
	span.Tag("lang_type", "java")
	span.End()


}

func many() {
	ctx := context.Background()
	span, _, _ := tracer.CreateLocalSpan(ctx, go2sky.WithOperationName("six"))
	span.Tag("lang_type", "java")
	span.End()
}

func second() {
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