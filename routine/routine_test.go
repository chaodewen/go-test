package main

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

func doStuff(wg *sync.WaitGroup, i int) {
	fmt.Printf("goroutine id %d\n", i)
	time.Sleep(300 * time.Second)
	fmt.Printf("goroutine id %d\n", i)
	wg.Done()
}

func TestRoutine(t *testing.T) {
	var wg sync.WaitGroup
	workers := 10

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go doStuff(&wg, i)
	}
	wg.Wait()
}
