package main

import (
	"github.com/astaxie/beego/logs"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	for i := 0; i < 10; i++ {
		/*
			e := email.NewEmail()
			e.From = "Jordan Wright <test@gmail.com>"
			e.To = []string{"test@example.com"}
			e.Bcc = []string{"test_bcc@example.com"}
			e.Cc = []string{"test_cc@example.com"}
			e.Subject = "Awesome Subject"
			e.Text = []byte("Text Body is, of course, supported!")
			e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

			//SMTP := os.Getenv("SMTP")
			//SMTPPORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

			err := e.Send("10.10.17.11:25", nil)
			fmt.Println(err)
		*/

		m := gomail.NewMessage()
		//m.SetHeader("From", "youngtip@gmail.com")
		m.SetHeader("From", "do-not-reply@naddic.com")
		m.SetHeader("To", "youngtip@gmail.com")
		//m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
		//m.SetHeader("Subject", "golang test")
		m.SetHeader("Subject", "title")
		m.SetBody("text/html", "test")
		//m.Attach("template.html") // attach whatever you want

		d := gomail.NewDialer("10.10.17.11", 25, "", "")

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			//panic(err)
			//beego.Error("send email error: ", err)
			logs.Error("send email error: ", err)
		} else {
			logs.Info("success send email to ")
		}
	}

}
