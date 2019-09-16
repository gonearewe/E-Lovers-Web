package email

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/go-gomail/gomail"
)

type email struct {
	from     string `addr of email sender`
	to       string `addr of email receiver`
	host     string `email host providing smtp service`
	password string `password of email sender`
	*gomail.Message
}

func newEmail(receiver string) *email {
	m := &email{
		from:     beego.AppConfig.String("emailsender"),
		to:       receiver,
		host:     beego.AppConfig.String("emailhost"),
		password: beego.AppConfig.String("emailsenderpasswd"),
		Message:  gomail.NewMessage(),
	}
	m.SetHeader("From", m.from)
	m.SetHeader("To", m.to)
	return m
}

func SendVerifyCodeEmail(targetEmail string) (verifiedCode string, err error) {

	//"cristmactavish@outlook.com"
	m := newEmail(targetEmail)
	m.SetHeader("Subject", "电子爱好者协会注册验证码")

	message := "<i>Hello ,你的验证码是：</i></hr><h3>CODE</h3></hr>"
	rand.Seed(time.Now().Unix())
	num := rand.Int63n(999999)
	code := fmt.Sprintf("%06d", num)
	strings.ReplaceAll(message, "CODE", code)
	m.SetBody("text/html", message)

	return code, m.send()
	//m.Attach("/home/Alex/lolcat.jpg")

	// d := gomail.NewDialer("", 587, "", "")

	// // Send the email to Bob, Cora and Dan.
	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("Send Successfully")
	// }
}
func (m *email) send() error {
	d := gomail.NewDialer(m.host, 587, m.from, m.password)
	return d.DialAndSend(m.Message)
}
