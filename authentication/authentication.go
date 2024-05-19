package authentication

import "email-app/src/entity"

func LoginUser(email string, password string, userList entity.USER_LIST) string {
	user := getUserByEmail(email, userList)

	if user == (entity.User{}) || !comparePassword(password, user.Password) {
		return "Email or password is incorrect"
	}

	if !user.IsVerified {
		return "User is not verified"
	}

	return "Login successful"
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

func RegisterUser(name string, email string, password string, userList *entity.USER_LIST) string {
	if getUserByEmail(email, *userList) != (entity.User{}) {
		return "Email already registered"
	}

	isFinished := false
	i := 0

	for !isFinished || i < len(userList) {
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

		i++
	}

	if !isFinished {
		return "User list is full"
	}

	return "Registration successful"
}

func LoginAsAdmin(email string, password string, adminList entity.USER_ADMIN_LIST) string {
	admin := getAdminByEmail(email, adminList)

	if admin == (entity.UserAdmin{}) || !comparePassword(password, admin.Password) {
		return "Email or password is incorrect"
	}

	return "Login successful"
}

func getAdminByEmail(email string, adminList entity.USER_ADMIN_LIST) entity.UserAdmin {
	for _, admin := range adminList {
		if admin.Email == email {
			return admin
		}
	}
	return entity.UserAdmin{}
}

func RegisterAdmin(name string, email string, password string, adminList *entity.USER_ADMIN_LIST) string {
	if getAdminByEmail(email, *adminList) != (entity.UserAdmin{}) {
		return "Email already registered"
	}

	isFinished := false
	i := 0

	for !isFinished || i < len(adminList) {
		if adminList[i] == (entity.UserAdmin{}) {
			adminList[i] = entity.UserAdmin{
				Id:       i + 1,
				Name:     name,
				Email:    email,
				Password: password,
			}
			isFinished = true
		}

		i++
	}

	if !isFinished {
		return "Admin list is full"
	}

	return "Registration successful"

}

func VerifyUser(email string, userList *entity.USER_LIST) string {
	user := getUserByEmail(email, *userList)

	if user == (entity.User{}) {
		return "User not found"
	}

	user.IsVerified = true

	return "User verified"
}
