package main

import (
	"encoding/base64"
	"testing"
	"fmt"
)

func TestMisc(t *testing.T) {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}

func TestTwo(t *testing.T) {
	decoded, err :=  base64.StdEncoding.DecodeString("YWJj")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decoded)
}