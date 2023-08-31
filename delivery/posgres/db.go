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

	defer rows.Close()
	fmt.Printf(" ________________________________________________________\n")
	fmt.Printf("|id          |name          |number          |description|\n")
	fmt.Printf("---------------------------------------------------------\n")
	for rows.Next() {
		var ar1, ar2, ar3, ar4 string
		rows.Scan(&ar1, &ar2, &ar3, &ar4)

		fmt.Printf("|%-12s |%-12s |%-12s |%-12s\n", ar1, ar2, ar3, ar4)
		fmt.Printf("---------------------------------------------------------\n")
	}

}

func DeleteUsers(connection *psx.Conn) {

	fmt.Printf("Delete by:\n 1. id\n 2. name\n 3. number\n 4. description\n")
	n := 0
	fmt.Scan(&n)
	tx, err := connection.Begin(context.TODO())
	defer tx.Commit(context.Background())
	if err != nil {
		log.Fatal("Cannot begin transaction", err)
	}

	if n == 1 {
		fmt.Print("Enter id: ")
		id := 0
		fmt.Scan(&id)
		_, err = tx.Exec(context.Background(), "DELETE FROM info WHERE id=$1", id)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		} else {
			fmt.Printf("Deleted\n")
		}
	} else if n == 2 {
		fmt.Print("Enter name: ")
		name := ""
		fmt.Scan(&name)
		_, err = tx.Exec(context.Background(), "DELETE FROM info WHERE name=$1", name)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		} else {
			fmt.Printf("Deleted\n")
		}
	} else if n == 3 {
		fmt.Print("Enter number: ")
		number := 0
		fmt.Scan(&number)
		_, err = tx.Exec(context.Background(), "DELETE FROM info WHERE number=$1", number)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		} else {
			fmt.Printf("Deleted\n")
		}
	} else if n == 4 {
		fmt.Print("Enter description: ")
		description := ""
		fmt.Scan(&description)
		_, err = tx.Exec(context.Background(), "DELETE FROM info WHERE description=$1", description)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		} else {
			fmt.Printf("Deleted\n")
		}
	} else {
		fmt.Printf("Incorrect column\n")
		return
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
		tx, err := connection.Begin(context.TODO())
		if err != nil {
			log.Fatal("Cannot begin transaction", err)
		}

		queryStr := "INSERT INTO info (name, number, description) VALUES ('" + name + "'" + "," + strconv.Itoa(number) + "," + "'" + descripton + "'" + ");"
		_, err = tx.Exec(context.Background(), queryStr)
		if err != nil {
			log.Fatal("Cannot execute query", err)
		}

		err = tx.Commit(context.Background())
		if err != nil {
			log.Fatal("Cannot commit transaction", err)
		}
	} else {
		return errors.New("incorrect data")
	}
	return nil
}
