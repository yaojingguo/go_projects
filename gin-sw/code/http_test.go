package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestBasics(t *testing.T) {
	{
		resp, err := http.Get("http://example.com/")
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("GET status: %s", resp.Status)
		}
		resp.Body.Close()
	}
	{
		resp, err := http.Post("https://httpbin.org/post", "application/json", strings.NewReader("hi"))
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("POST status: %s", resp.Status)
		}
		resp.Body.Close()
	}
	{
		resp, err := http.PostForm("https://httpbin.org/post",
			url.Values{"custname": {"yaojingguo"}})
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("POST FORM status: %s", resp.Status)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("body: %s", body)
		}
		resp.Body.Close()
	}
}

func TestAdvancedCustomization(t *testing.T) {
	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("status: %s", resp.Status)
	}
}

func TestBasicCustomization(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			t.Logf("req URL: %s",req.URL)
			t.Log("via:")
			for _, r := range via {
				t.Logf("  via URL: %s", r.URL)
			}
			return nil
		},
	}

	resp, err := client.Get("https://www.jd.com/xx")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("status: %s", resp.Status)
	}
	resp.Body.Close()
}
