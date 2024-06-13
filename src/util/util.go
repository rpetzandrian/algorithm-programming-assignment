package util

import (
	"email-app/src/entity"
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("clear") // Use "clear" command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CheckForExitInput[T string | int](input T, nextStep func()) {
	switch strInput := any(input).(type) {
	case string:
		if strings.ToLower(strInput) == "cancel" {
			nextStep()
		}
	case int:
		if strInput == -1 {
			nextStep()
		}
	}
}
func GenerateUserSeed(USERS *entity.USER_LIST) {
	USERS[0] = entity.User{Id: 1, Name: "test", Email: "test@test.com", Password: "12345", IsVerified: true}
	USERS[1] = entity.User{Id: 2, Name: "test2", Email: "test2@test.com", Password: "12345", IsVerified: true}
	USERS[2] = entity.User{Id: 2, Name: "test3", Email: "test3@test.com", Password: "12345", IsVerified: true}
}

func GenerateAdminSeed(ADMINS *entity.USER_ADMIN_LIST) {
	ADMINS[0] = entity.UserAdmin{Id: 1, Name: "admin", Email: "admin@test.com", Password: "12345"}
}

func GenerateEmailSeed(EMAILS *entity.EMAIL_LIST) {
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
}
