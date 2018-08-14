package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "mysql"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Person struct {
	ID   string `json:"ID"`
	Name string `json:"FIRSTNAME"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	//inserNewPerson()
	listAllPerson()

	//findPerson()

}

/*
func findPerson() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golang")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, firstname FROM person where id = ?", "1fe820b4-344b-4491-8e81-2b19cf5c0b22").Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(tag.ID)
	fmt.Println(tag.Name)

}

*/

func listAllPerson() {
	db := dbConn()

	results, err := db.Query("SELECT id, firstname FROM person")
	if err != nil {
		panic(err.Error())
	}

	person := Person{}
	res := []Person{}

	for results.Next() {

		var id string
		var firstname string

		err = results.Scan(&id, &firstname)
		if err != nil {
			panic(err.Error())
		}

		person.ID = id
		person.Name = firstname

		res = append(res, person)

	}

	personJson, _ := json.Marshal(res)
	err = ioutil.WriteFile("output.json", personJson, 0644)

	fmt.Println(res)
	defer db.Close()
}

func inserNewPerson() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golang")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("insert into person (id, firstname, lastname) values ('1fe820b4-344b-4491-8e81-2b19cf5c0b22', 'john', 'doe')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
