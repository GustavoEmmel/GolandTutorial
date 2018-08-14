package main

import "fmt"

type Person struct {
	Name string
	Number int
}

func main(){

	vault_dwellers := make(map[int]Person)
	// map[int] index type
	// Person content type

	for i := 1; i <= 5 ; i++ {
		var number int = i
		//fmt.Println(Person{"Vault Dweller nro", number});
		vault_dwellers[number] = Person{"Vault Dweller nro", number}
	}

	fmt.Println(vault_dwellers);
	fmt.Println(vault_dwellers[1]);

}