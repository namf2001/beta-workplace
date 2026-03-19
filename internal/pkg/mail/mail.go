package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strconv"

	"github.com/namf2001/beta-workplace/config"
)

// Mail configuration constants
const (
	// MIME and content type constants
	MIMEHeader = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Email subject constants
	SubjectRegistration   = "Subject: Beta Workplace - Xác thực tài khoản của bạn\n"
	SubjectForgotPassword = "Subject: Beta Workplace - Đặt lại mật khẩu\n"

	// Template paths
	TemplateRegistration   = "internal/pkg/mail/templates/template_forgot_password.html"
	TemplateForgotPassword = "internal/pkg/mail/templates/template_registration.html"
)

// SendMailRegistration sends registration verification email
func SendMailRegistration(name, email string, code string) error {
	cfg := config.GetConfig()
	auth := smtp.PlainAuth("", cfg.MailSMTPUser, cfg.MailSMTPPassword, cfg.MailSMTPHost)

	templateData := struct {
		Name string
		Code string
	}{
		name,
		code,
	}

	body, err := parseTemplate(TemplateRegistration, templateData)
	if err != nil {
		return err
	}

	msg := []byte(SubjectRegistration + MIMEHeader + "\n" + body)
	addr := cfg.MailSMTPHost + ":" + strconv.Itoa(cfg.MailSMTPPort)
	if err = smtp.SendMail(addr, auth, cfg.MailEmailFrom, []string{email}, msg); err != nil {
		return err
	}

	return nil
}

// SendMailForgotPassword sends forgot password verification email
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

	body, err := parseTemplate(TemplateForgotPassword, templateData)
	if err != nil {
		return err
	}

	msg := []byte(SubjectForgotPassword + MIMEHeader + "\n" + body)
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
