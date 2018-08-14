package main

import "fmt"

func main() {

	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	i := 0 
	for i<10{
		fmt.Println(i)
		i+=5
	}

	//infinti loop while true equivalent
	/*
	for {
		fmt.Println("Do stuff")
	}
	*/

	x := 5 
	for {
		fmt.Println("Do stuff", x)
		x+= 3
		if x > 25{
			break
		}
	}


}