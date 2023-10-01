package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}
	for i := 0; i < 4; i++ {
		mu.Lock()
		wg.Add(1)
		go printInterface(&wg, &mu, data1, i)

		mu.Lock()
		wg.Add(1)
		go printInterface(&wg, &mu, data2, i)
	}

	wg.Wait()
}

func printInterface(wg *sync.WaitGroup, mu *sync.Mutex, data []interface{}, i int) {
	defer mu.Unlock()
	defer wg.Done()
	fmt.Printf("%+v %d\n", data, i)
}
