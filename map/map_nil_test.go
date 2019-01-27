package _map

import (
	"testing"
	"fmt"
)

func TestMapNil(t *testing.T) {
	m := map[int]int {
		123: 123,
	}

	fmt.Printf("bool: %+v\n", m[1] == 0)
}
