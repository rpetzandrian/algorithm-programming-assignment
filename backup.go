// package menu

// import (
// 	"email-app/src/decorative"
// 	"email-app/src/entity"
// 	"email-app/src/util"
// 	"fmt"
// )

// var routes entity.USER_TYPE_LIST
// var userTypeIndex, routeIndex, choiceIndex int
// var USER_CRUD entity.USER_LIST
// var USER_ADMIN_CRUD entity.USER_ADMIN_LIST
// var EMAIL_CRUD entity.EMAIL_LIST

// func InitRoutes() {
// 	userTypeIndex = -1
// 	routeIndex = -1
// 	choiceIndex = -1

// 	routes[0] = entity.UserType{
// 		UserType: util.ADMIN_TYPE,
// 		RouteList: [10]entity.Route{
// 			{
// 				RouteName: util.ADMIN_AUTH_MENU,
// 				RouteFunc: func(choiceIndex *int) {
// 					decorative.HeaderTemplate()

// 					// Menambahkan menu user dan admin

// 					decorative.PrintLine()
// 					decorative.PrintTitle(" Admin Auth Menu ")
// 					decorative.PrintDecorativeLine()
// 					decorative.PrintMenu(1, "Register")
// 					decorative.PrintMenu(2, "Login")
// 					decorative.PrintMenu(3, "Back")
// 					decorative.PrintDecorativeLine()
// 					decorative.PrintInstruction(" Choose the number of the menu to continue ")
// 					decorative.PrintBottomLine()

// 					navigateChoiceIndex(choiceIndex)
// 				},
// 				ChoiceList: [4]entity.Choice{
// 					{
// 						ChoiceText: "Register",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
// 							navigateRoute(util.ADMIN_REGISTER_MENU, userTypeIndex, routeIndex, choiceIndex)
// 						},
// 					},
// 					{
// 						ChoiceText: "Login",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
// 							navigateRoute(util.ADMIN_LOGIN_MENU, userTypeIndex, routeIndex, choiceIndex)
// 						},
// 					},
// 					{
// 						ChoiceText: "Back",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 							navigateRoute(util.ADMIN_TYPE, userTypeIndex, routeIndex, choiceIndex)
// 						},
// 					},
// 				},
// 			},
// 			{
// 				RouteName: util.ADMIN_REGISTER_MENU,
// 				RouteFunc: func(choiceIndex *int) {
// 					fmt.Println("HELLO WORLD ADMIN REGISTER")
// 				},
// 			},
// 			{
// 				RouteName: util.ADMIN_LOGIN_MENU,
// 				RouteFunc: func(choiceIndex *int) {
// 					fmt.Println("HELLO WORLD ADMIN LOGIN")

// 				},
// 			},
// 			{
// 				RouteName: util.ADMIN_APPROVAL_MENU,
// 				RouteFunc: func(choiceIndex *int) {
// 					decorative.HeaderTemplate()
// 					// Menambahkan menu user dan admin
// 					decorative.PrintLine()
// 					decorative.PrintTitle(" Admin Approval and Rejection Menu ")
// 					decorative.PrintDecorativeLine()
// 					decorative.PrintMenu(1, "Approve/reject user")
// 					decorative.PrintMenu(2, "Back")
// 					decorative.PrintDecorativeLine()
// 					decorative.PrintInstruction(" Choose the number of the menu to continue ")
// 					decorative.PrintBottomLine()

// 					navigateChoiceIndex(choiceIndex)
// 				},
// 				ChoiceList: [4]entity.Choice{
// 					{
// 						ChoiceText: "Approve/reject user",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Back",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
// 							navigateRoute(util.ADMIN_LOGIN_MENU, userTypeIndex, routeIndex, choiceIndex)
// 						},
// 					},
// 				},
// 			},
// 			{
// 				RouteName: util.ADMIN_APPROVE_REJECT_MENU,
// 				RouteFunc: func(choiceIndex *int) {

// 				},
// 			},
// 		},
// 	}
// 	routes[1] = entity.UserType{
// 		UserType: util.USER_TYPE,
// 		RouteList: [10]entity.Route{
// 			{
// 				RouteName: util.USER_AUTH_MENU,
// 				RouteFunc: func(choiceIndex *int) {

// 				},
// 				ChoiceList: [4]entity.Choice{
// 					{
// 						ChoiceText: "Register",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Login",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Back",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 				},
// 			},
// 			{
// 				RouteName: util.USER_SUB_MENU,
// 				RouteFunc: func(choiceIndex *int) {

// 				},
// 				ChoiceList: [4]entity.Choice{
// 					{
// 						ChoiceText: "Send email",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Inbox",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Outbox",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 					{
// 						ChoiceText: "Back",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 				},
// 			},
// 			{
// 				RouteName: util.USER_REGISTER_MENU,
// 				RouteFunc: func(choiceIndex *int) {

// 				},
// 				ChoiceList: [4]entity.Choice{
// 					{
// 						ChoiceText: "Back",
// 						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func navigateRoute(name string, userTypeIndex *int, routeIndex *int, choiceIndex *int) {
// 	fmt.Println(name)
// 	if name == util.ADMIN_TYPE || name == util.USER_TYPE {
// 		if name == util.ADMIN_TYPE {
// 			*userTypeIndex = -1
// 			*routeIndex = -1
// 			*choiceIndex = -1
// 			Menu()
// 		}
// 		if name == util.USER_TYPE {
// 			*userTypeIndex = -1
// 			*routeIndex = -0
// 			*choiceIndex = -1
// 			Menu()
// 		}
// 	} else {
// 		for i := 0; i < len(routes); i++ {
// 			for j := 0; j < len(routes[i].RouteList); j++ {
// 				if routes[i].RouteList[j].RouteName == name {
// 					*userTypeIndex = i
// 					*routeIndex = j
// 					*choiceIndex = -1
// 					Menu()
// 					break
// 				}
// 			}
// 		}
// 	}
// }

// func navigateChoiceIndex(choiceIndex *int) {
// 	fmt.Scan(choiceIndex)
// 	*choiceIndex -= 1
// 	Menu()
// }

// func Menu() {
// 	fmt.Println(userTypeIndex, routeIndex, choiceIndex)
// 	if userTypeIndex != -1 && routeIndex != -1 {
// 		if choiceIndex != -1 {
// 			/* Jika user sudah memilih nomor input, maka akan menampilkan fungsi yang ada pada
// 			struct Choice yaitu berupa ChoiceFunc*/
// 			routes[userTypeIndex].RouteList[routeIndex].ChoiceList[choiceIndex].ChoiceFunc(&userTypeIndex, &routeIndex, &choiceIndex)
// 		} else {
// 			/* Jika user belum memilih nomor input, maka akan menampilkan fungsi yang ada pada
// 			struct Route yaitu RouteFunc untuk memberikan pilihan kepada user*/
// 			routes[userTypeIndex].RouteList[routeIndex].RouteFunc(&choiceIndex)
// 		}
// 	} else {
// 		PrintStartMenu(&userTypeIndex, &routeIndex)
// 	}
// }

// func PrintStartMenu(userTypeIndex *int, routeIndex *int) {

// 	decorative.HeaderTemplate()

// 	// Menambahkan menu user dan admin
// 	decorative.PrintLine()
// 	decorative.PrintTitle(" Primary Menu ")
// 	decorative.PrintDecorativeLine()
// 	decorative.PrintMenu(1, "Admin Menu")
// 	decorative.PrintMenu(2, "User Menu")
// 	decorative.PrintDecorativeLine()
// 	decorative.PrintInstruction(" Choose the number of the menu to continue ")
// 	decorative.PrintBottomLine()

// 	fmt.Scan(userTypeIndex)
// 	*userTypeIndex -= 1
// 	*routeIndex = 0
// 	Menu()
// 	// Menambahkan pesan penutup dengan warna yang berbeda
// 	// color.New(color.FgHiYellow, color.Bold).Println("\n🌟 Thanks for using this app! 🌟")
// }