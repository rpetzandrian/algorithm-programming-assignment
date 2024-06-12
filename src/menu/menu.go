package menu

import (
	"email-app/src/decorative"
	"email-app/src/entity"
	"email-app/src/features/authentication"
	"email-app/src/features/emails"
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
var selectedEmailIdx entity.Email

func InitRoutes() {

	// Inisialisasi Route
	userTypeIndex = -1 // 0 for admin 1 for user
	routeIndex = -1    // menu
	choiceIndex = -1   // sub menu

	// Inisialisasi data User
	USERS[0] = entity.User{Id: 1, Name: "test", Email: "test@test.com", Password: "12345", IsVerified: true}
	USERS[1] = entity.User{Id: 2, Name: "test2", Email: "test2@test.com", Password: "12345", IsVerified: true}
	USERS[2] = entity.User{Id: 2, Name: "test3", Email: "test3@test.com", Password: "12345", IsVerified: true}
	ADMINS[0] = entity.UserAdmin{Id: 1, Name: "admin", Email: "admin@test.com", Password: "12345"}
	EMAILS[0] = entity.Email{Id: 1, From: "test@test.com", To: "test2@test.com", IsRead: false, Subject: "test email 1", Body: "Ini test email saja. Jangan diubah dlu ya 1,.. hello world!", Timestamp: "2021-01-01 00:00:01"}
	EMAILS[1] = entity.Email{Id: 2, From: "test2@test.com", To: "test3@test.com", IsRead: false, Subject: "test email 2", Body: "Ini test email saja. Jangan diubah dlu ya 2,.. hello world!", Timestamp: "2021-01-01 00:00:02"}
	EMAILS[2] = entity.Email{Id: 3, From: "test2@test.com", To: "test@test.com", IsRead: false, Subject: "test email 3", Body: "Ini test email saja. Jangan diubah dlu ya 3,.. hello world!", Timestamp: "2021-01-01 00:00:03"}
	EMAILS[3] = entity.Email{Id: 4, From: "test@test.com", To: "test2@test.com", IsRead: false, Subject: "test email 4", Body: "Ini test email saja. Jangan diubah dlu ya 4,.. hello world!", Timestamp: "2021-01-01 00:00:04"}
	EMAILS[4] = entity.Email{Id: 5, From: "test@test.com", To: "test2@test.com", IsRead: false, Subject: "test email 5", Body: "Ini test email saja. Jangan diubah dlu ya 5,.. hello world!", Timestamp: "2021-01-01 00:00:05"}
	EMAILS[5] = entity.Email{Id: 6, From: "test2@test.com", To: "test@test.com", IsRead: false, Subject: "test email 6", Body: "Ini test email saja. Jangan diubah dlu ya 6,.. hello world!", Timestamp: "2021-01-01 00:00:06"}
	EMAILS[6] = entity.Email{Id: 7, From: "test3@test.com", To: "test2@test.com", IsRead: false, Subject: "test email 7", Body: "Ini test email saja. Jangan diubah dlu ya 7,.. hello world!", Timestamp: "2021-01-01 00:00:07"}
	EMAILS[7] = entity.Email{Id: 8, From: "test2@test.com", To: "test@test.com", IsRead: false, Subject: "test email 8", Body: "Ini test email saja. Jangan diubah dlu ya 8,.. hello world!", Timestamp: "2021-01-01 00:00:08"}
	EMAILS[8] = entity.Email{Id: 9, From: "test@test.com", To: "test4@test.com", IsRead: false, Subject: "test email 9", Body: "Ini test email saja. Jangan diubah dlu ya 9,.. hello world!", Timestamp: "2021-01-01 00:00:09"}
	EMAILS[9] = entity.Email{Id: 10, From: "test2@test.com", To: "test@test.com", IsRead: false, Subject: "test email 10", Body: "Ini test email saja. Jangan diubah dlu ya 10,.. hello world!", Timestamp: "2021-01-01 00:00:10"}
	EMAILS[10] = entity.Email{Id: 11, From: "test@test.com", To: "test2@test.com", IsRead: false, Subject: "test email 11", Body: "Ini test email saja. Jangan diubah dlu ya 11,.. hello world!", Timestamp: "2021-01-01 00:00:11"}
	EMAILS[11] = entity.Email{Id: 12, From: "test3@test.com", To: "test@test.com", IsRead: false, Subject: "test email 12", Body: "Ini test email saja. Jangan diubah dlu ya 12,.. hello world!", Timestamp: "2021-01-01 00:00:12"}

	CurrentLogged = entity.LoggedUser{Id: -1, Name: "", Email: "", Role: -1}

	routes[0] = entity.UserType{
		UserType: util.ADMIN_TYPE,
		RouteList: [10]entity.Route{
			{
				RouteName: util.ADMIN_AUTH_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {
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

					navigateInputIndex(3, choiceIndex)

				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: util.ADMIN_REGISTER_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage[string]("Admin Register Page")

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
									printStatus = util.PRINT_STATUS_SUCCESS
									printText = status
								}

								decorative.PrintStatus(printStatus, printText)
							}

							decorative.PrintText("Press any key to navigate to login page: ")
							var key int
							fmt.Scan(&key)
							navigateRoute(util.ADMIN_LOGIN_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
					{
						ChoiceText: util.ADMIN_LOGIN_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage[string]("Admin Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							for !loggedIn {
								decorative.ResetPrintStatus(&printStatus, &printText)
								email, password := authentication.InputUserLogin(func() {
									navigateRoute(util.ADMIN_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, message := authentication.LoginAsAdmin(email, password, ADMINS, &CurrentLogged)

								if err {
									printStatus = util.PRINT_STATUS_ERROR
									printText = message
								} else {
									loggedIn = true
								}

								decorative.PrintStatus(printStatus, printText)
							}

							navigateRoute(util.ADMIN_APPROVAL_MENU, userTypeIndex, routeIndex, choiceIndex)
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
				RouteName: util.ADMIN_APPROVAL_MENU,
				RouteFunc: func(printStatus *string, printText *string, choiceIndex *int) {
					HeaderTemplate()
					// Menambahkan menu user dan admin
					decorative.PrintLine()
					decorative.PrintTitle(" Admin Approval Menu ")
					decorative.PrintDecorativeLine()
					decorative.PrintMenu(1, "Approve user")
					decorative.PrintMenu(2, "Logout")
					decorative.PrintDecorativeLine()
					decorative.PrintInstruction(" Choose the number of the menu to continue ")
					decorative.PrintBottomLine()

					navigateInputIndex(2, choiceIndex)

				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceText: util.ADMIN_APPROVE_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderAdminMenu()
							headerPage[int]("Approve User Page")

							decorative.PrintInfo(" List Unverified User")

							id := 0
							for *userTypeIndex == 0 && *routeIndex == 1 && *choiceIndex == 0 {
								authentication.RetrieveUnverifiedUser(USERS)
								fmt.Println("Enter user id to approve: ")
								fmt.Scan(&id)

								util.CheckForExitInput[int](id, func() {
									navigateRoute(util.ADMIN_APPROVAL_MENU, userTypeIndex, routeIndex, choiceIndex)
								})

								err, message := authentication.VerifyUser(id, &USERS)

								if err {
									decorative.PrintAlert(message)
								} else {
									decorative.PrintInfo(message)
								}

							}
						},
					},
					{
						ChoiceText: util.LOGOUT,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.LogoutUser(&CurrentLogged)

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
							headerPage[string]("User Register Page")

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
									printStatus = util.PRINT_STATUS_SUCCESS
									printText = status
								}

								decorative.PrintStatus(printStatus, printText)
							}

							decorative.PrintText("input any key:")
							var key int
							fmt.Scan(&key)

							navigateRoute(util.USER_AUTH_LOGIN_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
					{
						ChoiceText: util.USER_AUTH_LOGIN_MENU,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.HeaderTemplate()
							headerPage[string]("User Login Page")
							isLoggedIn := false

							for !isLoggedIn {

								if printStatus == util.PRINT_STATUS_ERROR {
									decorative.PrintStatus(printStatus, printText)
								} else {
									decorative.PrintNothing()
								}

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
					decorative.PrintMenu(3, "Logout")
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
							HeaderUserMenu()
							headerPage[string]("Send Email Page")

							sent := false
							for !sent {
								decorative.ResetPrintStatus(&printStatus, &printText)
								to, subject, body := emails.WriteEmail(&CurrentLogged, func() {
									navigateRoute(util.USER_SUB_MENU, userTypeIndex, routeIndex, choiceIndex)
								})
								err, message := emails.SendEmail(CurrentLogged.Email, to, subject, body, &EMAILS)

								if err {
									printStatus = util.PRINT_STATUS_ERROR
									printText = message
								} else {
									printStatus = util.PRINT_STATUS_SUCCESS
									printText = message
									sent = true
								}

								decorative.PrintStatus(printStatus, printText)
							}

							navigateRoute(util.USER_SUB_MENU_SEND_EMAIL, userTypeIndex, routeIndex, choiceIndex)
						},
					},
					{
						ChoiceText: util.USER_SUB_MENU_INBOX,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							HeaderUserMenu()
							headerPage[int]("Inbox Page")

							mail := emails.RetrieveEmails(EMAILS, CurrentLogged.Email)
							totalIdx := emails.ShowEmailList(mail)

							idx := 0
							decorative.PrintInfo(" Input email number: ")
							inputsMenus(totalIdx, &idx)
							util.CheckForExitInput[int](idx, func() {
								navigateRoute(util.USER_SUB_MENU, userTypeIndex, routeIndex, choiceIndex)
							})

							selectedEmailIdx = mail[idx-1]

							*choiceIndex = 3
							Menu()
						},
					},
					{
						ChoiceText: util.LOGOUT,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.LogoutUser(&CurrentLogged)
							navigateRoute(util.USER_AUTH_MENU, userTypeIndex, routeIndex, choiceIndex)
						},
					},
					{
						ChoiceText: util.USER_SUB_MENU_EMAIL_LIST,
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							decorative.ResetPrintStatus(&printStatus, &printText)
							HeaderUserMenu()
							headerPage[int]("Email Page")

							fmt.Println("Selected Email:  ", selectedEmailIdx) // debug,.. need deleted
							list := emails.EmailList(selectedEmailIdx.From, selectedEmailIdx.To, EMAILS)
							emails.ShowEmailList(list)

							for *userTypeIndex == 1 && *routeIndex == 1 && *choiceIndex == 3 {
								var key int
								inputsMenus(0, &key)
								util.CheckForExitInput[int](key, func() {
									navigateRoute(util.USER_SUB_MENU_INBOX, userTypeIndex, routeIndex, choiceIndex)
								})
							}
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
		menuConfig := util.GetMenuConfig(name)
		*userTypeIndex = menuConfig.UserTypeIdx
		*routeIndex = menuConfig.RouteIdx
		*choiceIndex = menuConfig.ChoiceIdx
		Menu()
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

func SimpleHeaderTemplate() {
	decorative.PrintLine()
	decorative.PrintSubtitle(" Welcome to Email App")
	decorative.PrintBottomLine()
}

func HeaderAuthMenu(userType int) {
	SimpleHeaderTemplate()

	decorative.PrintLine()
	if userType == 0 {
		decorative.PrintTitle(" Admin Auth Menu ")
	} else {
		decorative.PrintTitle(" User Auth Menu ")
	}
	decorative.PrintBottomLine()

}

func HeaderUserMenu() {
	decorative.PrintLine()
	decorative.PrintSubtitle(" Welcome " + CurrentLogged.Name)
	decorative.PrintBottomLine()

	decorative.PrintLine()
	decorative.PrintTitle(" User Dashboard ")
	decorative.PrintBottomLine()
}

func HeaderAdminMenu() {
	decorative.PrintLine()
	decorative.PrintSubtitle(" Welcome " + CurrentLogged.Name)
	decorative.PrintBottomLine()

	decorative.PrintLine()
	decorative.PrintTitle(" Admin Dashboard ")
	decorative.PrintBottomLine()
}

func headerPage[T any](page string, opt ...bool) {
	decorative.PrintLine()
	decorative.PrintSubtitle(page)

	var input T

	if len(opt) < 1 || (len(opt) > 0 && !opt[0]) {
		decorative.PrintEmptyLine()
		switch any(input).(type) {
		case string:
			decorative.PrintInstruction("Type cancel and press enter to back....")
		case int:
			decorative.PrintInstruction("Type -1 and press enter to back....")
		}
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
	if *input == -1 {
		return false
	}

	if *input < 1 || *input > max {
		return true
	}

	return false
}
