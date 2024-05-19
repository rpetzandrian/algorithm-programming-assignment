package main

import (
	"email-app/src/entity"
	"email-app/src/menu"
	"fmt"
)

var USERS entity.USER_LIST
var USER_ADMINS entity.USER_ADMIN_LIST
var EMAILS entity.EMAIL_LIST

func main() {
	stateMenu := 0

	menu.PrintStartMenu()

	fmt.Print("Enter the number of the menu:  ")
	fmt.Scan(&stateMenu)
}
