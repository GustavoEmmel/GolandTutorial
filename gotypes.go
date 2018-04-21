package main

import ("fmt")

func add(x float64, y float64) float64 {
	return x + y
}

func add2(x, y float64) float64 {
	return x + y
}

func multiple(a,b string) (string, string) {
	return a, b
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(add(num1, num2))

	// num3, num4 := 5.6, 9.5
	// fmt.Println(add2(num3, num4))

	w1, w2 := "hey", "there"

	fmt.Println(multiple(w1, w2))

/*
	var a int = 64
	var b float64 = float64(a)

	x : = a // x will be type int
*/

}