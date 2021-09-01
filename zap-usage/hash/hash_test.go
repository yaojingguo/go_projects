package hash

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))
}