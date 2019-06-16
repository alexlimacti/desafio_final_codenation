package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")

	if err != nil {
		err = fmt.Errorf("error open db: %v", err)
	}
	schema := `CREATE TABLE place (
		country text,
		city text NULL,
		telcode integer);`
	result, err := db.Exec(schema)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result ", result)
}
