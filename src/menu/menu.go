package menu

import (
	"email-app/src/decorative"
	"email-app/src/entity"
	"email-app/src/features/authentication"
	"email-app/src/util"
	"fmt"
)

var routes entity.USER_TYPE_LIST
var userTypeIndex, routeIndex, choiceIndex int

var USERS entity.USER_LIST
var ADMINS entity.USER_ADMIN_LIST
var EMAILS entity.EMAIL_LIST
var CurrentLogged entity.LoggedUser

func InitRoutes() {
	// Inisialisasi Route
	userTypeIndex = -1 // 0 for admin 1 for user
	routeIndex = -1    // menu
	choiceIndex = -1   // sub menu

	// Inisialisasi data User
	CurrentLogged = entity.LoggedUser{Id: -1, Name: "", Email: "", Role: -1}

	routes[0] = entity.UserType{
		UserType: "Admin",
		RouteList: [10]entity.Route{
			{
				RouteId:   0,
				RouteName: "Admin Auth Menu",
				RouteFunc: func(choiceIndex *int) {
					HeaderTemplate()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" Admin Auth Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Register")
					decorative.PrintMenu(2, "Login")
					decorative.PrintMenu(3, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					inputsMenus(3, choiceIndex)

					*choiceIndex -= 1
					// Pilihan input nomor dari user dikurang 1 dan kita memanggil Menu untuk mengubah dan memanggil fungsi yang ada pada struct Choice
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Register",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderTemplate()
							headerPage("Admin Register Page")

							// Function to register user
							loggedIn := CurrentLogged.Id != -1 && CurrentLogged.Role == 1 // check is user loggedin

							for !loggedIn {
								name, email, password := authentication.InputUserRegister()
								err, message := authentication.RegisterAdmin(name, email, password, &ADMINS)

								if err {
									fmt.Println(message)
								} else {
									fmt.Println(message)
									loggedIn = true
								}
							}

							*choiceIndex = -1
							*routeIndex = 1

							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Login",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderTemplate()
							headerPage("Admin Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							for !loggedIn {
								email, password := authentication.InputUserLogin()
								err, message := authentication.LoginAsAdmin(email, password, ADMINS, &CurrentLogged)

								if err {
									fmt.Println(message)
								} else {
									fmt.Println(message)
									loggedIn = true
								}
							}

							*choiceIndex = -1
							*routeIndex = 1

							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 3,
						ChoiceText:   "Back",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							*userTypeIndex = -1
							*routeIndex = -1
							*choiceIndex = -1
							Menu()
						},
					},
				},
			},
			{
				RouteId:   1,
				RouteName: "Admin Approval and Rejection Menu",
				RouteFunc: func(choiceIndex *int) {
					HeaderTemplate()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" Admin Approval and Rejection Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Approve/reject user")
					decorative.PrintMenu(2, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					inputsMenus(2, choiceIndex)

					*choiceIndex -= 1
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Approve/reject user",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.RetrieveUnverifiedUser(USERS)

							*choiceIndex = -1
							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Back",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							*userTypeIndex = 0
							*routeIndex = 0
							*choiceIndex = -1
							Menu()
						},
					},
				},
			},
		},
	}
	routes[1] = entity.UserType{
		UserType: "User",
		RouteList: [10]entity.Route{
			{
				RouteId:   0,
				RouteName: "User Auth Menu",
				RouteFunc: func(choiceIndex *int) {
					HeaderTemplate()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" User Auth Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Register")
					decorative.PrintMenu(2, "Login")
					decorative.PrintMenu(3, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					inputsMenus(3, choiceIndex)

					*choiceIndex -= 1
					// Pilihan input nomor dari user dikurang 1 dan kita memanggil Menu untuk mengubah dan memanggil fungsi yang ada pada struct Choice
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Register",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderTemplate()
							headerPage("User Register Page")

							// Function to register user
							loggedIn := CurrentLogged.Id != -1 && CurrentLogged.Role == 1 // check is user loggedin

							if loggedIn {
								*choiceIndex = -1
								*routeIndex = 1

								util.ClearScreen()
								Menu()
							}

							finishRegister := false
							for !loggedIn && !finishRegister {
								name, email, password := authentication.InputUserRegister()
								err, message := authentication.RegisterUser(name, email, password, &USERS)

								if err {
									fmt.Println(message)
								} else {
									fmt.Println(message)
									finishRegister = true
								}
							}

							fmt.Println("Waiting For Admin to approve Register Request")
							fmt.Print("Press any key to Login Page")
							var key string
							fmt.Scan(&key)

							*choiceIndex = 1
							*routeIndex = 0

							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Login",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderTemplate()
							headerPage("User Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							errorMessage := ""
							for !loggedIn {
								decorative.PrintAlert(errorMessage)
								decorative.PrintInfo(" Input instruction: ")

								var key int
								inputsMenus(2, &key)
								if key == 2 {
									*routeIndex = 0
									*choiceIndex = -1

									util.ClearScreen()
									Menu()
								}

								email, password := authentication.InputUserLogin()
								err, message := authentication.LoginUser(email, password, USERS, &CurrentLogged)

								if err {
									errorMessage = message
								} else {
									loggedIn = true
								}
							}

							*choiceIndex = -1
							*routeIndex = 1

							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 0,
						ChoiceText:   "Back",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
				},
			},
			{
				RouteId:   1,
				RouteName: "User Sub Menu",
				RouteFunc: func(choiceIndex *int) {
					HeaderUserMenu()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" User Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Send Email")
					decorative.PrintMenu(2, "Inbox")
					decorative.PrintMenu(3, "Outbox")
					decorative.PrintMenu(4, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					inputsMenus(3, choiceIndex)

					*choiceIndex -= 1
					// Pilihan input nomor dari user dikurang 1 dan kita memanggil Menu untuk mengubah dan memanggil fungsi yang ada pada struct Choice
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Send email",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Inbox",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceNumber: 3,
						ChoiceText:   "Outbox",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceNumber: 0,
						ChoiceText:   "Back",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
				},
			},
		},
	}
}

func Menu() {
	fmt.Printf("user type : %d route index : %d choice index : %d\n", userTypeIndex, routeIndex, choiceIndex) // guide
	if userTypeIndex != -1 && routeIndex != -1 {
		if choiceIndex != -1 {
			/* Jika user sudah memilih nomor input, maka akan menampilkan fungsi yang ada pada
			struct Choice yaitu berupa ChoiceFunc*/
			routes[userTypeIndex].RouteList[routeIndex].ChoiceList[choiceIndex].ChoiceFunc(&userTypeIndex, &routeIndex, &choiceIndex)
		} else {
			/* Jika user belum memilih nomor input, maka akan menampilkan fungsi yang ada pada
			struct Route yaitu RouteFunc untuk memberikan pilihan kepada user*/
			routes[userTypeIndex].RouteList[routeIndex].RouteFunc(&choiceIndex)
		}
	} else {
		PrintStartMenu(&userTypeIndex, &routeIndex)
	}
}

func PrintStartMenu(userTypeIndex *int, routeIndex *int) {

	HeaderTemplate()

	// Menambahkan menu user dan admin
	decorative.PrintLine()
	decorative.PrintTitle(" Primary Menu ")
	decorative.PrintDecorativeLine()
	decorative.PrintMenu(1, "Admin Menu")
	decorative.PrintMenu(2, "User Menu")
	decorative.PrintDecorativeLine()
	decorative.PrintInstruction(" Choose the number of the menu to continue ")
	decorative.PrintBottomLine()

	inputsMenus(2, userTypeIndex)

	*userTypeIndex -= 1
	*routeIndex = 0
	util.ClearScreen()
	Menu()
	// Menambahkan pesan penutup dengan warna yang berbeda
	// color.New(color.FgHiYellow, color.Bold).Println("\n🌟 Thanks for using this app! 🌟")
}
func HeaderTemplate() {
	// Mencetak tampilan dengan dekorasi dan informasi proyek yang lebih menarik
	decorative.PrintLine()
	decorative.PrintTitle(" Alpro Assignment ")
	decorative.PrintDecorativeLine()
	decorative.PrintSubtitle(" EMAIL APP ")
	decorative.PrintEmptyLine()
	decorative.PrintSubtitle(" Created by: ")
	decorative.PrintAuthor(" Rico x Daffa ")
	decorative.PrintBottomLine()
}

func HeaderUserMenu() {
	decorative.PrintLine()
	decorative.PrintSubtitle(" Welcome " + CurrentLogged.Name)
	decorative.PrintBottomLine()
}

func headerPage(page string) {
	decorative.PrintLine()
	decorative.PrintSubtitle(page)
	decorative.PrintEmptyLine()
	decorative.PrintInstruction(" Input number to continue:  ")
	decorative.PrintInstruction(" Press 1 to continue ")
	decorative.PrintInstruction(" Press 2 to back")
	decorative.PrintBottomLine()
}

func inputsMenus(menuMax int, input *int) {
	fmt.Scan(input)

	check := true
	for check {
		err := validateMenuInputs(menuMax, input)

		if err {
			fmt.Println("Enter correct input number: ")
			fmt.Scan(input)
		} else {
			check = false
		}
	}
}

func validateMenuInputs(max int, input *int) (err bool) {
	if *input < 1 || *input > max {
		return true
	}

	return false
}

func infoPage(info string) {
	decorative.PrintLine()
	decorative.PrintSubtitle(info)
	decorative.PrintBottomLine()
}
