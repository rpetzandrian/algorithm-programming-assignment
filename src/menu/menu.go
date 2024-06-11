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
var printStatus string
var printText string

func InitRoutes() {
	// Inisialisasi Route
	userTypeIndex = -1 // 0 for admin 1 for user
	routeIndex = -1    // menu
	choiceIndex = -1   // sub menu

	// Inisialisasi data User
	CurrentLogged = entity.LoggedUser{Id: -1, Name: "", Email: "", Role: -1}

	routes[0] = entity.UserType{
		UserType: util.ADMIN_TYPE,
		RouteList: [10]entity.Route{
			{
				RouteName: util.ADMIN_AUTH_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {
					decorative.HeaderTemplate()
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

					navigateInputIndex(3, choiceIndex)

				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: util.ADMIN_REGISTER_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage("Admin Register Page")

							// Function to register user
							loggedIn := CurrentLogged.Id != -1 && CurrentLogged.Role == 1 // check is user loggedin

							for !loggedIn {
								name, email, password := authentication.InputUserRegister(func() {
									navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, status := authentication.RegisterAdmin(name, email, password, &ADMINS)

								if err {
									printStatus = util.PRINT_STATUS_ERROR
									printText = status
									navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								} else {
									navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								}
							}
						},
					},
					{
						ChoiceText: util.ADMIN_LOGIN_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage("Admin Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							for !loggedIn {
								email, password := authentication.InputUserLogin(func() {
									navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, message := authentication.LoginAsAdmin(email, password, ADMINS, &CurrentLogged)

								if err {
									fmt.Println(message)
								} else {
									fmt.Println(message)
									loggedIn = true
								}
							}

							navigateRoute(util.ADMIN_APPROVE_REJECT_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
					{
						ChoiceText: util.ADMIN_AUTH_MENU_BACK,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							navigateRoute(util.ADMIN_TYPE, userTypeIndex, routeIndex, choiceIndex)
						},
					},
				},
			},
			{
				RouteName: util.ADMIN_REGISTER_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {

				},
			},
			{
				RouteName: util.ADMIN_LOGIN_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {

				},
			},
			{
				RouteName: util.ADMIN_APPROVE_REJECT_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {
					decorative.HeaderTemplate()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" Admin Approval and Rejection Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Approve/reject user")
					decorative.PrintMenu(2, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					navigateInputIndex(2, choiceIndex)

				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: "Approve/reject user",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.RetrieveUnverifiedUser(USERS)

						},
					},
					{
						ChoiceText: "Back",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
				},
			},
		},
	}
	routes[1] = entity.UserType{
		UserType: util.USER_TYPE,
		RouteList: [10]entity.Route{
			{
				RouteName: util.USER_AUTH_MENU,
				RouteFunc: func(printStatus *string, printText *string, hoiceIndex *int) {
					decorative.HeaderTemplate()
					// Menambahkan menu user dan admin
					if *printStatus == util.PRINT_STATUS_ERROR {
						decorative.PrintStatus(*printStatus, *printText)
					} else if *printStatus == util.PRINT_STATUS_NOTHING {
						decorative.PrintNothing()
					}
					decorative.PrintLine()
					decorative.PrintTitle(" User Auth Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Register")
					decorative.PrintMenu(2, "Login")
					decorative.PrintMenu(3, "Back")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					navigateInputIndex(3, &choiceIndex)
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: util.USER_AUTH_REGISTER_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage("User Register Page")

							finishRegister := false

							for !finishRegister {
								name, email, password := authentication.InputUserRegister(func() {
									decorative.ResetPrintStatus(&printStatus, &printText)
									navigateRoute(util.USER_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, status := authentication.RegisterUser(name, email, password, &USERS)

								if err {
									printStatus = util.PRINT_STATUS_ERROR
									printText = status
									navigateRoute(util.USER_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								} else {
									finishRegister = true
									navigateRoute(util.USER_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								}
								fmt.Println(&printStatus)
							}
						},
					},
					{
						ChoiceText: util.USER_AUTH_LOGIN_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

							isLoggedIn := false

							for !isLoggedIn {

								decorative.HeaderTemplate()

								if printStatus == util.PRINT_STATUS_ERROR {
									decorative.PrintStatus(printStatus, printText)
								} else {
									decorative.PrintNothing()
								}

								headerPage("User Login Page")

								email, password := authentication.InputUserLogin(func() {
									navigateRoute(util.USER_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, status := authentication.LoginUser(email, password, USERS, &CurrentLogged)

								printText = status

								if err {
									printStatus = util.PRINT_STATUS_ERROR
									util.ClearScreen()
								} else {
									isLoggedIn = true
									printStatus = util.PRINT_STATUS_SUCCESS
									navigateRoute(util.USER_SUB_MENU, userTypeIndex, routeIndex, choiceIndex)
								}
							}
						},
					},
					{
						ChoiceText: util.USER_AUTH_BACK_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							navigateRoute(util.ADMIN_TYPE, userTypeIndex, routeIndex, choiceIndex)
						},
					},
				},
			},
			{
				RouteName: util.USER_SUB_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {
					HeaderUserMenu()
					// Menambahkan menu user dan admin
					if *printStatus == util.PRINT_STATUS_SUCCESS {
						decorative.PrintStatus(*printStatus, *printText)
					}
					decorative.PrintLine()
					decorative.PrintTitle(" User Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Send Email")
					decorative.PrintMenu(2, "Inbox")
					decorative.PrintMenu(3, "Outbox")
					decorative.PrintMenu(4, "Log out")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					navigateInputIndex(3, choiceIndex)

					*choiceIndex -= 1
					// Pilihan input nomor dari user dikurang 1 dan kita memanggil Menu untuk mengubah dan memanggil fungsi yang ada pada struct Choice
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: util.USER_SUB_MENU_SEND_EMAIL,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceText: util.USER_SUB_MENU_INBOX,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceText: util.USER_SUB_MENU_OUTBOX,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceText: util.USER_SUB_MENU_LOGOUT,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							navigateRoute(util.USER_SUB_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
				},
			},
		},
	}
}

func navigateRoute(name string, userTypeIndex *int, routeIndex *int, choiceIndex *int) {
	if name == util.ADMIN_TYPE || name == util.USER_TYPE {
		if name == util.ADMIN_TYPE {
			*userTypeIndex = -1
			*routeIndex = -1
			*choiceIndex = -1
			Menu()
		}
		if name == util.USER_TYPE {
			*userTypeIndex = -1
			*routeIndex = -0
			*choiceIndex = -1
			Menu()
		}
	} else {
		for i := 0; i < len(routes); i++ {
			for j := 0; j < len(routes[i].RouteList); j++ {
				if routes[i].RouteList[j].RouteName == name {
					*userTypeIndex = i
					*routeIndex = j
					*choiceIndex = -1
					Menu()
					break
				}
			}
		}
	}
}

func Menu() {
	util.ClearScreen()
	if userTypeIndex != -1 && routeIndex != -1 {
		if choiceIndex != -1 {
			/* Jika user sudah memilih nomor input, maka akan menampilkan fungsi yang ada pada
			struct Choice yaitu berupa ChoiceFunc*/
			routes[userTypeIndex].RouteList[routeIndex].ChoiceList[choiceIndex].ChoiceFunc(&userTypeIndex, &routeIndex, &choiceIndex)
		} else {
			/* Jika user belum memilih nomor input, maka akan menampilkan fungsi yang ada pada
			struct Route yaitu RouteFunc untuk memberikan pilihan kepada user*/
			routes[userTypeIndex].RouteList[routeIndex].RouteFunc(&printStatus, &printText, &choiceIndex)
		}
	} else {
		PrintStartMenu(&userTypeIndex, &routeIndex)
	}
}

func PrintStartMenu(userTypeIndex *int, routeIndex *int) {

	decorative.HeaderTemplate()

	// Menambahkan menu user dan admin
	decorative.PrintLine()
	decorative.PrintTitle(" Primary Menu ")
	decorative.PrintDecorativeLine()
	decorative.PrintMenu(1, "Admin Menu")
	decorative.PrintMenu(2, "User Menu")
	decorative.PrintDecorativeLine()
	decorative.PrintInstruction(" Choose the number of the menu to continue ")
	decorative.PrintBottomLine()

	*routeIndex = 0
	navigateInputIndex(2, userTypeIndex)

	// Menambahkan pesan penutup dengan warna yang berbeda
	// color.New(color.FgHiYellow, color.Bold).Println("\n🌟 Thanks for using this app! 🌟")
}

func HeaderUserMenu() {
	decorative.PrintLine()
	decorative.PrintSubtitle(" Welcome " + CurrentLogged.Name)
	decorative.PrintBottomLine()
}

func headerPage(page string, opt ...bool) {

	decorative.PrintLine()
	decorative.PrintSubtitle(page)

	if len(opt) < 1 || (len(opt) > 0 && !opt[0]) {
		decorative.PrintEmptyLine()
		decorative.PrintInstruction("Type cancel and press enter to back....")
	}
	decorative.PrintBottomLine()
}

func navigateInputIndex(menuMax int, inputIndex *int) {
	valid := false

	for !valid {
		fmt.Print("Enter input number: ")
		fmt.Scan(inputIndex)
		status := validateChoiceInput(menuMax, inputIndex)

		if status {
			valid = true
		}
	}
	fmt.Println(*inputIndex)
	*inputIndex -= 1
	Menu()
}
func validateChoiceInput(max int, input *int) (status bool) {
	if *input >= 1 && *input <= max {
		return true
	}

	return false
}
