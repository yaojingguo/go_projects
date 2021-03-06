package kv

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3/concurrency"
	"testing"
	"time"
)

func TestMutexBasics(t *testing.T) {
	cli := NewClient()
	defer cli.Close()

	// Create two separate sessions for lock competition
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	defer s1.Close()
	pfx := "/TestMutexBasics"
	m1 := NewMutex(s1, pfx)

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()
	m2 := NewMutex(s2, pfx)

	ctx1 := context.Background()
	ctx2 := context.Background()

	// acquire lock for s1
	if err := m1.Lock(ctx1); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("acquired lock for s1 with key %s and rev %d\n", m1.myKey, m1.myRev)

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		if err := m2.Lock(ctx2); err != nil {
			t.Fatal(err)
		}
	}()

	var seconds time.Duration = 3
	fmt.Printf("sleeping for %d seconds\n", seconds)
	time.Sleep(seconds * time.Second)
	if err := m1.Unlock(ctx1); err != nil {
		t.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Printf("acquired lock for s2 with key %s and rev %d\n", m2.myKey, m2.myRev)
	if err := m2.Unlock(ctx2); err != nil {
		t.Fatal(err)
	}
}
