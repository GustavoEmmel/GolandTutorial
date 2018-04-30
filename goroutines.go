package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}
}

func main() {

	// go before a func start a lightwheit thread
	go say("hey")
	go say("there")

	// need sleep to see the print, otherwise the app will close and print nothing
	time.Sleep(time.Microsecond * 9900)

}
