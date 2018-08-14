package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {

	// defer works like first input last output
	defer wg.Done()
	for i := 0; i <= 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}

}

func main() {

	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()

}
