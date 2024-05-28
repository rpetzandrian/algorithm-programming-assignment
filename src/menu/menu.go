package menu

import (
	"email-app/src/decorative"
	"email-app/src/entity"
	"fmt"
)

var routes entity.USER_TYPE_LIST
var userTypeIndex, routeIndex, choiceIndex int

func InitRoutes() {
	userTypeIndex = -1
	routeIndex = -1
	choiceIndex = -1
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
					fmt.Scan(choiceIndex)
					*choiceIndex -= 1
					// Pilihan input nomor dari user dikurang 1 dan kita memanggil Menu untuk mengubah dan memanggil fungsi yang ada pada struct Choice
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Register",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

							// Memberi nilai 0 pada userTypeIndex
							*userTypeIndex = 0
							// Memberi nilai 1 pada routeIndex untuk memanggil fungsi yang ada pada userTypeIndex 0 dan routeIndex 1
							*routeIndex = 1
							// Memberi nilai -1 pada choiceIndex, -1 diberikan karena user belum melakukan input pada halaman yang dituju
							*choiceIndex = -1
							Menu()
						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Login",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {
							*userTypeIndex = 0
							*routeIndex = 1
							*choiceIndex = -1
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
					fmt.Scan(choiceIndex)
					*choiceIndex -= 1
					Menu()
				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Approve/reject user",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

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

				},
				ChoiceList: [4]entity.Choice{
					{
						ChoiceNumber: 1,
						ChoiceText:   "Register",
						ChoiceFunc: func(userTypeIndex *int, routeIndex *int, choiceIndex *int) {

						},
					},
					{
						ChoiceNumber: 2,
						ChoiceText:   "Login",
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
			{
				RouteId:   1,
				RouteName: "User Sub Menu",
				RouteFunc: func(choiceIndex *int) {

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
	fmt.Printf("user type : %d route index : %d choice index : %d\n", userTypeIndex, routeIndex, choiceIndex)
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

	fmt.Scan(userTypeIndex)
	*userTypeIndex -= 1
	*routeIndex = 0
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
