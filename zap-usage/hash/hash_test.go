package hash

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	data := []byte("The quick brown fox jumps over the lazy dog")
	hashBytes := md5.Sum(data)
	fmt.Println(len(hashBytes))
	hashStr := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashStr)
}