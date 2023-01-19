package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "eddi:eddi@123@tcp(localhost:3306)/golang_db")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var name string
	err = db.QueryRow("SELECT email FROM user WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	print("success")
}
