package sendto

import (
	"bytes"
	"fmt"
	"go-ecomm/global"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
	"strings"
)

const (
	SMTPHost     = "live.smtp.mailtrap.io"
	SMTPPort     = "587"
	SMTPUserName = "smtp@mailtrap.io"
	SMTPPassword = "6e4e5e52fb75879bf84c8464840a246b"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func Send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: from},
		To:      to,
		Subject: "OTP verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)
	fmt.Println(messageMail)
	auth := smtp.PlainAuth("", SMTPUserName, SMTPPassword, SMTPHost)
	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, to, []byte(messageMail))

	if err != nil {
		global.Logger.Error("Send Email Failed::", zap.Error(err))
		return err
	}

	return nil
}

func SendTemplateEmailOTp(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}
	return Send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		global.Logger.Error("Execute Template Failed::", zap.Error(err))
	}
	fmt.Println(htmlTemplate.String())
	return htmlTemplate.String(), nil
}

func send() {

}
