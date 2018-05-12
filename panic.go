package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recoverd in cleanup", r)
	}
	wg.Done()
}

func say(s string) {

	defer cleanup()
	for i := 0; i <= 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)

		if i == 2 {
			panic("Oh dear, a 2")
		}
	}

}

func main() {

	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()

}
