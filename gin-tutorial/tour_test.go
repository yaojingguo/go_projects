package main

import (
	"fmt"
	"testing"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func TestOne(t *testing.T) {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func TestTwo(t *testing.T) {
	for i, v := range pow {
		fmt.Printf("2**%d == %d\n", i, v)
	}
}
