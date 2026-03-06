package mail

import (
	"bytes"
	"html/template"

	"net/smtp"
	"strconv"

	"github.com/namf2001/beta-workplace/config"
)

// SendMailForgotPassword sends an email to the user when they forgot password
func SendMailForgotPassword(name, email string, code string) error {
	cfg := config.GetConfig()
	auth := smtp.PlainAuth("", cfg.MailSMTPUser, cfg.MailSMTPPassword, cfg.MailSMTPHost)

	templateData := struct {
		Name string
		Code string
	}{
		name,
		code,
	}

	body, err := parseTemplate("helper/mail/template_mail.html", templateData)
	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: Location Tracker forgot password!\n"
	msg := []byte(subject + mime + "\n" + body)
	addr := cfg.MailSMTPHost + ":" + strconv.Itoa(cfg.MailSMTPPort)
	if err = smtp.SendMail(addr, auth, cfg.MailEmailFrom, []string{email}, msg); err != nil {
		return err
	}

	return nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
