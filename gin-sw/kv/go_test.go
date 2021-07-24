package kv

import (
	"context"
	"testing"
)

type Person struct {
	Age int
	Name string
}

func TestLang(t *testing.T) {
	// ctx := context.Background()
	// if me, has := ctx.Value("abc").(Person); has {
	// 	t.Logf("person: %#v", me)
	// }
	// t.Log("no person in context")]
	work(nil, t)
}

func work(ctx context.Context, t *testing.T) {
	if me, has := ctx.Value("abc").(Person); has {
		t.Logf("person: %#v", me)
	}
	t.Log("no person in context")
}
