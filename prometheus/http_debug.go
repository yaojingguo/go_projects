package main

import (
	"fmt"
	"golang.org/x/net/trace"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func fooHandler(w http.ResponseWriter, req *http.Request) {
	tr := trace.New("mypkg.Foo", req.URL.Path)
	defer tr.Finish()
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
	tr.LazyPrintf("event %d happened", 99)
	////...
	//if err := somethingImportant(); err != nil {
	//	tr.LazyPrintf("somethingImportant failed: %v", err)
	//	tr.SetError()
	//}
}

func work() {
	for i := 1; i <= 5; i++ {
		tr := trace.New("mypkg.Foo", "background")
		defer tr.Finish()
		tr.LazyPrintf("working %d", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
		tr.LazyPrintf("done %d", i)
	}
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellow world")
}

type Fetcher struct {
	domain string
	events trace.EventLog
}

func NewFetcher(domain string) *Fetcher {
	return &Fetcher{
		domain,
		trace.NewEventLog("mypkg.Fetcher", domain),
	}
}

func (f *Fetcher) Fetch(path string) (string, error) {
	time.Sleep(time.Second * time.Duration(rand.Int31n(2)))
	resp, err := http.Get("http://" + f.domain + "/" + path)
	if err != nil {
		f.events.Errorf("Get(%q) = %v", path, err)
		return "", err
	}
	f.events.Printf("Get(%q) = %s", path, resp.Status)
	return resp.Status, nil
}

func (f *Fetcher) Close() error {
	f.events.Finish()
	return nil
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)

	go work()

	go func() {
		f := NewFetcher("oa.xdf.cn")
		for i := 0; i < 100; i++ {
			fmt.Println(f.Fetch("app/Meeting/list.php"))
		}
		f.Close()
	} ()

	fmt.Println("starting HTTP server")
	log.Fatal(http.ListenAndServe(":5555", nil))
}

