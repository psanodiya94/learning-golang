package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=Rishirich11!")
	if err != nil {
		log.Fatal("unable to connect ", err)
	}

	defer conn.Close()

	log.Println("connected to database!")

	err = conn.Ping()
	if err != nil {
		log.Fatal("can't ping database ", err)
	}

	log.Println("Pingged Databse!")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	insertQuery := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(insertQuery, "jash", "tiwani")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("inserted a row!")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	updateQuery := `update users set first_name = $1, last_name = $2 where id = $3`
	_, err = conn.Exec(updateQuery, "pooja", "rai", 6)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("row updated!")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	selectQuery := `select id, first_name, last_name from users where id = $1`
	var firstName, lastName string
	var id int

	row := conn.QueryRow(selectQuery, 2)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("one row fetched by id!")
	log.Println("Query Result : ", id, firstName, lastName)

	deleteQuery := `delete from users where id = $1`
	_, err = conn.Exec(deleteQuery, 7)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("row deleted!")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println("Record is ", id, firstName, lastName)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("--------------------------------------------------")

	return nil
}
