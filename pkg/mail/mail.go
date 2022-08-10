package mail

import (
	"log"
	"net/smtp"
)

func SendNews(body string) {
	from := "devhome.courses@gmail.com"
	password := "fgientufdzqtbehz"
	to := []string{
		"annguyenthe@vccorp.vn",
	}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject:New Posts!\n"

	go func() {
		message := []byte(subject + body)
		auth := smtp.PlainAuth("", from, password, host)
		err := smtp.SendMail(address, auth, from, to, message)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Email Sent Successfully!")
	}()
}
