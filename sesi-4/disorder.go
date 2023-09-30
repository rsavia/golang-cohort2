package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go printInterface(&wg, data1, i)
	}

	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go printInterface(&wg, data2, i)
	}

	wg.Wait()
}

func printInterface(wg *sync.WaitGroup, data []interface{}, i int) {
	defer wg.Done()
	fmt.Printf("%+v %d\n", data, i)
}

