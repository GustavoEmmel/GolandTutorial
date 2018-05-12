package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	// the channel block the goroutines
	v1 := <-fooVal
	v2 := <-fooVal

	fmt.Println(v1, v2)
}
