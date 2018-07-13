package main

import (
	"entity/person"
	"hello"
	"config/database"

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func dbConn() (db *sql.DB) {
	dbDriver := database.MYSQL_DRIVER
	dbUser := database.MYSQL_USER
	dbPass := database.MYSQL_PASS
	dbName := database.MYSQL_DBNAME
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	t := hello.Hello()

	fmt.Println(t)

	listAllPerson()

	StartService("8090")

}

func StartService(port string) {

	http.HandleFunc("/person", Person)
	//http.HandleFunc("/getuser", userGetHandler)
	http.ListenAndServe(":" + port, nil)
}

func Person(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET PERSON")
		case "PUT":
			fmt.Fprintf(w, "PUT PERSON")
		default:
			http.Error(w, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
	}
}

func listAllPerson() {
	db := dbConn()

	results, err := db.Query("SELECT id, firstname FROM person")
	if err != nil {
		panic(err.Error())
	}

	per := person.Person{}
	res := []person.Person{}

	for results.Next() {

		var id string
		var firstname string

		err = results.Scan(&id, &firstname)
		if err != nil {
			panic(err.Error())
		}

		per.ID = id
		per.Name = firstname

		res = append(res, per)

	}

	personJson, _ := json.Marshal(res)
	err = ioutil.WriteFile("output.json", personJson, 0644)

	fmt.Println(res)
	defer db.Close()
}
