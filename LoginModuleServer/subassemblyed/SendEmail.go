package subassemblyed

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"net/smtp"
	"strings"
	"time"
)

type Email struct {
	Address     string
	Password    string
	Name        string
	TAddress    string
	TName       string
	CcAddress   string
	CcName      string
	BccAddress  string
	BccName     string
	SuValue     string
	ContentType string
	ContentBody string
	FileName    string
	Host        string
	Port        string
}

func SendEmail(ContentBody string, TAddress string, TName string, SuValue string, args ...string) time.Time {
	UserEmail := "571500549@qq.com"
	Mail_Smtp_Port := ":25"
	Mail_Password := "qrgykefttqambbbe"
	Mail_Smtp_Host := "smtp.qq.com"
	auth := smtp.PlainAuth("", UserEmail, Mail_Password, Mail_Smtp_Host)
	to := []string{TAddress}
	nickname := TName
	user := UserEmail
	subject := SuValue
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := ContentBody
	var buf bytes.Buffer
	buf.WriteString("To: ")
	buf.WriteString(strings.Join(to, ","))
	buf.WriteString("\r\nFrom: ")
	buf.WriteString(nickname)
	buf.WriteString("<")
	buf.WriteString(user)
	buf.WriteString(">\r\nSubject: ")
	buf.WriteString(subject)
	buf.WriteString("\r\n")
	buf.WriteString(content_type)
	buf.WriteString("\r\n\r\n")
	buf.WriteString(body)
	msg := buf.Bytes()
	err := smtp.SendMail(Mail_Smtp_Host+Mail_Smtp_Port, auth, user, to, msg)
	if err != nil {
		logs.Warn("send mail error: %v", err)
	}
	now := time.Now()
	return now
}
