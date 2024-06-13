package emails

import (
	//"bufio"
	"email-app/src/decorative"
	"email-app/src/entity"
	"email-app/src/util"
	"fmt"

	//"os"
	"time"
)

func WriteEmail(currentUser *entity.LoggedUser, nextStep func()) (to, subject, body string) {
	//reader := bufio.NewReader(os.Stdin)

	fmt.Println("To :")
	fmt.Scan(&to)
	util.CheckForExitInput[string](to, nextStep)

	fmt.Println("Subject :")
	fmt.Scan(&subject)
	//subject, _ = reader.ReadString('\n')
	util.CheckForExitInput[string](subject, nextStep)

	fmt.Println("Body :")
	fmt.Scan(&body)
	//body, _ = reader.ReadString('\n')
	util.CheckForExitInput[string](body, nextStep)

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

func ReadEmail(email1, email2 string, emails *entity.EMAIL_LIST, currentLogin entity.LoggedUser) {
	var fromEmail, toEmail string
	if email1 != currentLogin.Email {
		fromEmail = email1
		toEmail = email2
	} else {
		fromEmail = email2
		toEmail = email1
	}

	for idx, email := range emails {
		if email.From == fromEmail && email.To == toEmail {
			emails[idx].IsRead = true
		}
	}
}

func RetrieveEmails(emails entity.EMAIL_LIST, email string) (result entity.EMAIL_LIST) {
	// get related email
	var emailList entity.EMAIL_LIST
	for i := 0; i < len(emails); i++ {
		if emails[i].To == email || emails[i].From == email {
			emailList[i] = emails[i]
		}
	}

	// sort email list by timestamp
	sortedEmailList := sortEmailByTimestamp(emailList)

	var user = make(map[string]bool)
	idx := 0
	for i := 0; i < len(sortedEmailList); i++ {
		if sortedEmailList[i] != (entity.Email{}) {
			if sortedEmailList[i].From == email {
				if !user[sortedEmailList[i].To] {
					result[idx] = sortedEmailList[i]
					user[sortedEmailList[i].To] = true
					idx++
				}
			} else if sortedEmailList[i].To == email {
				if !user[sortedEmailList[i].From] {
					result[idx] = sortedEmailList[i]
					user[sortedEmailList[i].From] = true
					idx++
				}
			}
		}
	}

	return
}

func EmailList(email1, email2 string, emails entity.EMAIL_LIST) (mails entity.EMAIL_LIST) {
	var emailList entity.EMAIL_LIST
	for i := 0; i < len(emails); i++ {
		if emails[i].To == email1 && emails[i].From == email2 {
			emailList[i] = emails[i]
		} else if emails[i].To == email2 && emails[i].From == email1 {
			emailList[i] = emails[i]
		}
	}

	mails = reassignEmail(sortEmailByTimestamp(emailList))

	return
}

func reassignEmail(emails entity.EMAIL_LIST) (mails entity.EMAIL_LIST) {
	x := 0
	for i := 0; i < len(emails); i++ {
		if emails[i] != (entity.Email{}) {
			emails[x] = emails[i]
			x += 1

			emails[i] = entity.Email{}
		}
	}

	return emails
}

func sortEmailByTimestamp(emails entity.EMAIL_LIST) (sortedEmailList entity.EMAIL_LIST) {
	sortedEmailList = emails
	for i := 1; i < len(sortedEmailList); i++ {
		key := sortedEmailList[i]
		j := i - 1
		for j >= 0 && sortedEmailList[j].Timestamp > key.Timestamp {
			sortedEmailList[j+1] = sortedEmailList[j]
			j = j - 1
		}
		sortedEmailList[j+1] = key
	}

	return
}

func ShowEmailList(emails entity.EMAIL_LIST) (counter int) {
	fmt.Println("==============================================")
	for i := 0; i < len(emails); i++ {
		if emails[i] != (entity.Email{}) {
			decorative.PrintInfo(fmt.Sprintf("No: %d :: %20s\n", i+1, emails[i].Timestamp))
			decorative.PrintText(fmt.Sprintf("From: %s", emails[i].From))
			decorative.PrintText(fmt.Sprintf("To: %s", emails[i].To))
			decorative.PrintWarning(fmt.Sprintf("Subject: %s\n", emails[i].Subject))
			decorative.PrintText(fmt.Sprintf("Body: %s", emails[i].Body))

			if condition := emails[i].IsRead; condition {
				decorative.PrintInfo("R")
			} else {
				decorative.PrintAlert("U")

			}
			fmt.Println("==============================================")

			counter++
		}
	}
	decorative.PrintInfo("End of List.")
	fmt.Println("==============================================")

	return
}

func DeleteEmail(emails *entity.EMAIL_LIST, emailId int) (err bool, message string) {
	err = true
	message = "Email not found"
	for idx, email := range emails {
		if email.Id == emailId {
			emails[idx] = entity.Email{}
			err = false
			message = "Email deleted successfully"
		}
	}

	return
}
