// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kv

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/client/v3/concurrency"
	"testing"
	"time"
	"log"

	spb "go.etcd.io/etcd/api/v3/mvccpb"
	v3 "go.etcd.io/etcd/client/v3"
)

var (
	ErrKeyExists      = errors.New("key already exists")
	ErrWaitMismatch   = errors.New("unexpected wait result")
	ErrTooManyClients = errors.New("too many clients")
	ErrNoWatcher      = errors.New("no watcher channel")
)
var (
	endPoints   = []string{"localhost:2379"}
	dialTimeout = 1 * time.Second

)

func NewClient() *v3.Client {
	cli, err := v3.New(v3.Config{
		Endpoints:   endPoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func NewSession(t *testing.T) (cli *v3.Client, session *concurrency.Session) {
	cli, err := v3.New(v3.Config{
		Endpoints:   endPoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		t.Fatal(err)
	}
	session, err = concurrency.NewSession(cli)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// deleteRevKey deletes a key by revision, returning false if key is missing
func deleteRevKey(kv v3.KV, key string, rev int64) (bool, error) {
	cmp := v3.Compare(v3.ModRevision(key), "=", rev)
	req := v3.OpDelete(key)
	txnresp, err := kv.Txn(context.TODO()).If(cmp).Then(req).Commit()
	if err != nil {
		return false, err
	} else if !txnresp.Succeeded {
		return false, nil
	}
	return true, nil
}

func claimFirstKey(kv v3.KV, kvs []*spb.KeyValue) (*spb.KeyValue, error) {
	for _, k := range kvs {
		ok, err := deleteRevKey(kv, string(k.Key), k.ModRevision)
		if err != nil {
			return nil, err
		} else if ok {
			return k, nil
		}
	}
	return nil, nil
}

func Get(context context.Context, cli *v3.Client, key string) ([]byte, error) {
	defer cli.Close()

	response, err := cli.Get(context, key)
	if err != nil {
		return nil, err
	}
	if response.Count != 1 {
		return nil, fmt.Errorf("return %d key values", response.Count)
	}
	return response.Kvs[0].Value, nil
}

