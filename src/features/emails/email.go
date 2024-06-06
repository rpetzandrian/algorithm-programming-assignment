package emails

import (
	"bufio"
	"email-app/src/entity"
	"fmt"
	"os"
	"time"
)

func WriteEmail(currentUser *entity.LoggedUser) (to, subject, body string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("To :")
	fmt.Scanln(&to)

	fmt.Println("Subject :")
	subject, _ = reader.ReadString('\n')

	fmt.Println("Body :")
	body, _ = reader.ReadString('\n')

	return
}

func SendEmail(from string, to string, subject string, body string, emails *entity.EMAIL_LIST) (err bool, message string) {
	isFinished := false
	i := 0

	for !isFinished {
		if emails[i] == (entity.Email{}) {
			emails[i] = entity.Email{
				Id:        i + 1,
				From:      from,
				To:        to,
				Subject:   subject,
				Body:      body,
				IsRead:    false,
				Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			}
			isFinished = true
		}

		if i == len(emails)-1 {
			isFinished = true
			return true, "Email list is full"
		}

		i++
	}

	return false, "Email sent successfully"
}

func ReadEmail(fromEmail, toEmail string, emails *entity.EMAIL_LIST) {
	for idx, email := range emails {
		if email.From == fromEmail && email.To == toEmail {
			emails[idx].IsRead = true
		}
	}
}
