package authentication

import (
	"email-app/src/entity"
	"fmt"
)

func LoginUser(email string, password string, userList entity.USER_LIST, currUser *entity.LoggedUser) (err bool, message string) {
	user := getUserByEmail(email, userList)

	if user == (entity.User{}) || !comparePassword(password, user.Password) {
		return true, "Email or password is incorrect"
	}

	if !user.IsVerified {
		return true, "User is not verified"
	}

	currUser.Id = user.Id
	currUser.Name = user.Name
	currUser.Email = user.Email
	currUser.Role = 1

	return false, "Login successful"
}

func getUserByEmail(email string, userList entity.USER_LIST) entity.User {
	for _, user := range userList {
		if user.Email == email && user.IsVerified { // Access the email field using the correct syntax
			return user
		}
	}
	return entity.User{}
}

func comparePassword(password string, userPassword string) bool {
	return password == userPassword
}

func RegisterUser(name string, email string, password string, userList *entity.USER_LIST) (err bool, message string) {
	if getUserByEmail(email, *userList) != (entity.User{}) {
		return true, "Email already registered"
	}

	isFinished := false
	i := 0

	for !isFinished {
		if userList[i] == (entity.User{}) {
			userList[i] = entity.User{
				Id:         i + 1,
				Name:       name,
				Email:      email,
				Password:   password,
				IsVerified: false,
			}
			isFinished = true
		}

		if i == len(userList)-1 {
			isFinished = true
			return true, "User list is full"
		}

		i++
	}

	return false, "Registration successful"
}

func LoginAsAdmin(email string, password string, adminList entity.USER_ADMIN_LIST, currUser *entity.LoggedUser) (err bool, message string) {
	admin := getAdminByEmail(email, adminList)

	if admin == (entity.UserAdmin{}) || !comparePassword(password, admin.Password) {
		return true, "Email or password is incorrect"
	}

	currUser.Id = admin.Id
	currUser.Name = admin.Name
	currUser.Email = admin.Email
	currUser.Role = 0

	return false, "Login successful"
}

func getAdminByEmail(email string, adminList entity.USER_ADMIN_LIST) entity.UserAdmin {
	for _, admin := range adminList {
		if admin.Email == email {
			return admin
		}
	}
	return entity.UserAdmin{}
}

func RegisterAdmin(name string, email string, password string, adminList *entity.USER_ADMIN_LIST) (err bool, message string) {
	if getAdminByEmail(email, *adminList) != (entity.UserAdmin{}) {
		return true, "Email already registered"
	}

	isFinished := false
	i := 0

	for !isFinished {
		if adminList[i] == (entity.UserAdmin{}) {
			adminList[i] = entity.UserAdmin{
				Id:       i + 1,
				Name:     name,
				Email:    email,
				Password: password,
			}
			isFinished = true
		}

		if i == len(adminList)-1 {
			isFinished = true
			return true, "Admin list is full"
		}

		i++
	}

	return false, "Registration successful"
}

func VerifyUser(id int, userList *entity.USER_LIST) (err bool, message string) {
	user, idx := getUserById(id, *userList)

	if user == (entity.User{}) {
		return true, "User not found"
	}

	userList[idx].IsVerified = true

	return false, "User verified successfully"
}

func InputUserRegister() (name, email, password string) {
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your password: ")
	fmt.Scan(&password)

	return
}

func InputUserLogin() (email, password string) {
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your password: ")
	fmt.Scan(&password)

	return
}

func LogoutUser(currUser *entity.LoggedUser) {
	currUser.Id = -1
	currUser.Name = ""
	currUser.Email = ""
	currUser.Role = -1
}

func getUnverifiedUserList(userList entity.USER_LIST) (unverifiedUserList entity.USER_LIST) {
	i := 0
	for _, user := range userList {
		if !user.IsVerified {
			unverifiedUserList[i] = user
			i++
		}
	}

	return
}

func RetrieveUnverifiedUser(userList entity.USER_LIST) {
	unverifiedUserList := getUnverifiedUserList(userList)

	len := 0
	for _, user := range unverifiedUserList {
		if user.Id != 0 {
			len++
			fmt.Printf("UserId %d | Name: %s | Email: %s\n", user.Id, user.Name, user.Email)
		}
	}

	if len == 0 {
		fmt.Println("No unverified user")
	}
}

func getUserById(id int, userList entity.USER_LIST) (entity.User, int) {
	for idx, user := range userList {
		if user.Id == id {
			return user, idx
		}
	}

	return entity.User{}, -1
}
