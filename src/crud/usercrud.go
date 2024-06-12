package crud

import (
	"bufio"
	"email-app/src/decorative"
	"email-app/src/entity"
	"fmt"
	"os"
	"strings"
)

func RegisterUser(userCrud *entity.USER_LIST) {
	var name, email string
	decorative.HeaderTemplate()
	decorative.PrintLine()
	decorative.PrintTitle(" Admin Auth Menu ")
	decorative.PrintDecorativeLine()
	decorative.PrintMenu(1, "Enter your name")
	decorative.PrintMenu(2, "Enter your email")
	decorative.PrintMenu(3, "Enter your password")
	decorative.PrintDecorativeLine()
	decorative.PrintMenu(4, "After that you have to wait for approve/reject by admin")
	decorative.PrintDecorativeLine()
	decorative.PrintInstruction(" Enter your name : ")
	fmt.Scan(&name)
	decorative.PrintBottomLine()
	decorative.PrintInstruction(" Enter your email : ")
	fmt.Scan(&email)
	decorative.PrintBottomLine()
	reader := bufio.NewReader(os.Stdin)
	decorative.PrintInstruction(" Enter your password : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split the input into words
	words := strings.Fields(input)

	// Replace each word with asterisks
	for i := range words {
		words[i] = strings.Repeat("*", len(words[i]))
	}

	// Join the words back into a single string
	output := strings.Join(words, " ")
	fmt.Println(output)
	decorative.PrintBottomLine()
}
