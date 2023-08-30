package posgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	psx "github.com/jackc/pgx/v5"
)

const DB_URL = "postgresql://postgres:telBook@localhost:5432/telBook"

func ConnectDB() *psx.Conn {

	config, _ := psx.ParseConfig(DB_URL)

	conn, err := psx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Cannot to connect to database", err)
		return nil
	}

	return conn

}
func ShowUsers(connection *psx.Conn) {
	rows, err := connection.Query(context.Background(), "SELECT * FROM info;")

	if err != nil {
		log.Fatal("Cannot execute query", err)
	}

	for rows.Next() {
		var ar1, ar2, ar3, ar4 string
		rows.Scan(&ar1, &ar2, &ar3, &ar4)
		fmt.Printf("%s %s %s %s\n", ar1, ar2, ar3, ar4)
	}
}

func DeleteUsers(connection *psx.Conn) {
	rows, err := connection.Query(context.Background(), "SELECT * FROM info;")

	if err != nil {
		log.Fatal("Cannot execute query", err)
	}

	for rows.Next() {
		var ar1, ar2, ar3, ar4 string
		rows.Scan(&ar1, &ar2, &ar3, &ar4)
		fmt.Printf("%s %s %s %s\n", ar1, ar2, ar3, ar4)
	}
}

func AddNewUsers(connection *psx.Conn) error {
	name := ""
	number := 0
	descripton := ""
	check := true
	fmt.Print("Enter name: ")
	if n, err := fmt.Scan(&name); n > 255 || err != nil {
		check = false
	}
	fmt.Print("Enter number: ")
	if _, err := fmt.Scan(&number); err != nil {
		check = false
	}
	fmt.Print("Enter description: ")
	if _, err := fmt.Scan(&descripton); err != nil {
		check = false
	}

	if check {
		queryStr := "INSERT INTO info (name, number, description) VALUES ('" + name + "'" + "," + strconv.Itoa(number) + "," + "'" + descripton + "'" + ");"
		_, err := connection.Query(context.Background(), queryStr)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		}
	} else {
		return errors.New("incorrect data")
	}
	return nil
}
