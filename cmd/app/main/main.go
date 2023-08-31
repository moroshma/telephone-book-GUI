package main

import (
	"cmd/app/main/main.go/delivery/posgres"
	"fmt"
)

func main() {
	con := posgres.ConnectDB()
	fl := true

	for fl {
		fmt.Print("\033[H\033[2J")
		printMenu()
		commd := 0
		fmt.Scan(&commd)

		if commd == 1 {
			posgres.ShowUsers(con)
		} else if commd == 2 {
			posgres.AddNewUsers(con)
		} else if commd == 3 {
			posgres.DeleteUsers(con)
		} else if commd == 4 {
			fl = false
		}
		fmt.Print("Press any key to continue...\n")
		fmt.Scanln()
	}
}

func printMenu() {
	fmt.Printf("1. Show all users\n")
	fmt.Printf("2. Add new user\n")
	fmt.Printf("3. Delete user\n")
	fmt.Printf("4. Exit\n")
}
