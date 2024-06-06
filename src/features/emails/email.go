package emails

import (
	"bufio"
	"email-app/src/decorative"
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

func RetrieveEmails(emails *entity.EMAIL_LIST, email string) (result entity.EMAIL_LIST) {
	// get related email
	var emailList entity.EMAIL_LIST
	for i := 0; i < len(emails); i++ {
		if emails[i].To == email || emails[i].From == email {
			emailList[i] = emails[i]
		}
	}

	// sort email list by timestamp
	sortedEmailList := emailList
	for i := 1; i < len(sortedEmailList); i++ {
		j := i
		for j > 0 && sortedEmailList[j-1].Timestamp < sortedEmailList[j].Timestamp {
			sortedEmailList[j-1], sortedEmailList[j] = sortedEmailList[j], sortedEmailList[j-1]
			j--
		}
	}

	var user = make(map[string]bool)
	for i := 0; i < len(sortedEmailList); i++ {
		if sortedEmailList[i].From == email {
			if !user[sortedEmailList[i].To] {
				result[i] = sortedEmailList[i]
				user[sortedEmailList[i].To] = true
			}
		} else {
			if !user[sortedEmailList[i].From] {
				result[i] = sortedEmailList[i]
				user[sortedEmailList[i].From] = true
			}
		}
	}

	return
}

func ShowEmailList(emails entity.EMAIL_LIST) (counter int) {
	fmt.Println("==============================================")
	for i := 0; i < len(emails); i++ {
		if emails[i] != (entity.Email{}) {
			decorative.PrintInfo(fmt.Sprintf("No: %d ;; From: %s    To: %s", counter+1, emails[i].From, emails[i].To))
			decorative.PrintWarning(fmt.Sprintf("Subject: %s", emails[i].Subject))
			decorative.PrintText(fmt.Sprintf("Body: %s", emails[i].Body))
			fmt.Println("==============================================")

			counter++
		}
	}
	decorative.PrintInfo("End of List. Input email number to see the detail")
	fmt.Println("==============================================")

	return
}
