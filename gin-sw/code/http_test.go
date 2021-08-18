package main

import (
	"net/http"
	"testing"
)

func TestBasics(t *testing.T) {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("status: %d", resp.Status)
	}
}
