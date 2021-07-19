package kv


import (
	"context"
	"testing"
)

//
func TestClientGet(t *testing.T) {
	cli := NewClient()
	defer cli.Close()

	val, err := Get(context.Background(), cli,"TestClientGet")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("value: %s", val)
}