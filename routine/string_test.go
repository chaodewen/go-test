package main

import (
	"fmt"
	"testing"
)

var (
	str = "abc"
)

func TestStringReadWrite(t *testing.T) {
	for i := 0; i < 1000; i++ {
		read()
		write(i)
		read()
	}
}

func read() {
	fmt.Println(str)
}

func write(i int) {
	str = fmt.Sprintf("%d: %s", i, str)
}
