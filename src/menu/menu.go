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
var selectedEmailIdx entity.Email

func InitRoutes() {

	// Inisialisasi Route
	userTypeIndex = -1 // 0 for admin 1 for user
	routeIndex = -1    // menu
	choiceIndex = -1   // sub menu

	// Inisialisasi data User
	USERS[0] = entity.User{Id: 1, Name: "test", Email: "test@test.com", Password: "12345", IsVerified: true}
	USERS[1] = entity.User{Id: 2, Name: "test2", Email: "test2@test.com", Password: "12345", IsVerified: true}
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
		UserType: "Admin",
		RouteList: [10]entity.Route{
			{
				RouteId:   0,
				RouteName: "Admin Auth Menu",
				RouteFunc: func(choiceIndex *int) {
					SimpleHeaderTemplate()
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
							HeaderAuthMenu(*userTypeIndex)
							headerPage("Admin Register Page")

							// Function to register user
							loggedIn := CurrentLogged.Id != -1 && CurrentLogged.Role == 1 // check is user loggedin

							errorMessage := ""
							for !loggedIn && *userTypeIndex == 0 && *routeIndex == 0 && *choiceIndex == 0 {
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
							HeaderAuthMenu(*userTypeIndex)
							headerPage("Admin Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							errorMessage := ""
							for !loggedIn && *routeIndex == 0 && *choiceIndex == 1 && *userTypeIndex == 0 {
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
				RouteName: "Admin Approval Menu",
				RouteFunc: func(choiceIndex *int) {
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

					inputsMenus(2, choiceIndex)

					*choiceIndex -= 1
					util.ClearScreen()
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Approve User",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderAdminMenu()
							headerPage("Approve User Page", true)

							decorative.PrintWarning(" Input 0 to back: ")
							decorative.PrintInfo(" List Unverified User")

							id := 0
							for *userTypeIndex == 0 && *routeIndex == 1 && *choiceIndex == 0 {
								authentication.RetrieveUnverifiedUser(USERS)
								fmt.Println("Enter user id to approve: ")
								fmt.Scan(&id)

								if id != 0 {
									err, message := authentication.VerifyUser(id, &USERS)

									if err {
										decorative.PrintAlert(message)
									} else {
										decorative.PrintInfo(message)
									}
								} else {
									*choiceIndex = -1
									util.ClearScreen()
									Menu()
								}
							}
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Logout",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.LogoutUser(&CurrentLogged)

							*userTypeIndex = 0
							*routeIndex = 0
							*choiceIndex = -1

							util.ClearScreen()
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
					SimpleHeaderTemplate()
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
							HeaderAuthMenu(*userTypeIndex)
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
							errorMessage := ""
							for !loggedIn && !finishRegister && *userTypeIndex == 1 && *routeIndex == 0 && *choiceIndex == 0 {
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
							HeaderAuthMenu(*userTypeIndex)
							headerPage("User Login Page")

							// Function to login user
							loggedIn := CurrentLogged.Id != -1
							errorMessage := ""
							for !loggedIn && *routeIndex == 0 && *choiceIndex == 1 && *userTypeIndex == 1 {
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
							*choiceIndex = -1
							*routeIndex = -1

							util.ClearScreen()
							Menu()
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
					decorative.PrintMenu(3, "Logout")
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
							HeaderUserMenu()
							headerPage("Send Email Page")

							errMessage := ""
							successMessage := ""
							for *userTypeIndex == 1 && *routeIndex == 1 && *choiceIndex == 0 {
								decorative.PrintAlert(errMessage)
								if successMessage != "" {
									infoPage(successMessage)
								}

								decorative.PrintInfo(" Input instruction: ")

								var key int
								inputsMenus(2, &key)
								if key == 2 {
									*routeIndex = 1
									*choiceIndex = -1

									util.ClearScreen()
									Menu()
								}

								to, subject, body := emails.WriteEmail(&CurrentLogged)
								err, message := emails.SendEmail(CurrentLogged.Email, to, subject, body, &EMAILS)

								if err {
									errMessage = message
								} else {
									successMessage = message
									util.ClearScreen()
									Menu()
								}
							}
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Inbox",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderUserMenu()
							headerPage("Inbox Page", true)

							mail := emails.RetrieveEmails(EMAILS, CurrentLogged.Email)
							totalIdx := emails.ShowEmailList(mail)

							idx := 0
							decorative.PrintWarning(fmt.Sprintf(" Input %d to back: ", totalIdx+1))
							decorative.PrintInfo(" Input email number: ")
							inputsMenus(totalIdx+1, &idx)

							if idx != totalIdx+1 {
								selectedEmailIdx = mail[idx-1]

								*choiceIndex = 3
								util.ClearScreen()
								Menu()
							} else {
								*choiceIndex = -1
								util.ClearScreen()
								Menu()
							}
						},
					},
					{
						ChoiceNumber: 3,
						ChoiceText:   "Logout",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							authentication.LogoutUser(&CurrentLogged)

							*userTypeIndex = 1
							*routeIndex = 0
							*choiceIndex = -1

							util.ClearScreen()
							Menu()
						},
					},
					{
						ChoiceNumber: 4,
						ChoiceText:   "Email List",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							HeaderUserMenu()
							headerPage("Email Page", true)

							fmt.Println("Selected Email:  ", selectedEmailIdx)
							list := emails.EmailList(selectedEmailIdx.From, selectedEmailIdx.To, EMAILS)
							emails.ShowEmailList(list)

							for *userTypeIndex == 1 && *routeIndex == 1 && *choiceIndex == 3 {
								decorative.PrintInfo(" Input 1 to back: ")

								var key int
								inputsMenus(1, &key)
								if key == 1 {
									selectedEmailIdx = entity.Email{}
									*choiceIndex = 1
									util.ClearScreen()
									Menu()
								}
							}
						},
					},
				},
			},
		},
	}
}

func Menu(opts ...interface{}) {
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

func headerPage(page string, opt ...bool) {
	decorative.PrintLine()
	decorative.PrintSubtitle(page)

	if len(opt) < 1 || (len(opt) > 0 && !opt[0]) {
		decorative.PrintEmptyLine()
		decorative.PrintInstruction(" Input number to continue:  ")
		decorative.PrintInstruction(" Press 1 to continue ")
		decorative.PrintInstruction(" Press 2 to back")
	}
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
