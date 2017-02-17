package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var workers = 100 // what if we have 1, 2, 25?

	if len(os.Args) == 2 {
		workers, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Printf("Worker: %d\n", workers)

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
