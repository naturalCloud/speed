package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/smtp"
	"strings"
)

type Mail struct {
	to         []string //接受者
	subject    string   //标题
	content    string   //邮件内容
	mailClient *smtp.Client
	host       string
	port       int16
	auth       smtp.Auth
	from       string
	pwd        string
	mailUser   string
	msg        []byte //构建完成的邮件字节
	context    *gin.Context
}

func NewMail(host, pwd, mailUser string, port int) (*Mail, error) {

	mail := &Mail{
		host:     host,
		port:     int16(port),
		mailUser: mailUser,
		pwd:      pwd,
	}

	var err error
	mail.mailClient, err = mail.sslClient()

	if err != nil {
		return nil, errors.New("mail client init error : " + err.Error())
	}

	mail.buildAuth()
	return mail, nil
}

//发送者
func (m *Mail) From(from string) *Mail {
	m.from = from + `<` + m.mailUser + `>`
	return m
}

//make  a smtp client
func (m *Mail) sslClient() (*smtp.Client, error) {
	addr := fmt.Sprintf("%s:%d", m.host, m.port)
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//构建auth
func (m *Mail) buildAuth() {
	m.auth = smtp.PlainAuth(
		"",
		m.mailUser,
		m.pwd,
		m.host,
	)
}

//设置发信人
func (m *Mail) To(user []string) *Mail {
	m.to = user
	return m
}

func (m *Mail) Subject(subject string) *Mail {
	m.subject = subject
	return m
}

func (m *Mail) Content(msg string) *Mail {
	m.content = msg
	return m
}

//设置消息
func (m *Mail) buildMsg() {

	header := map[string]string{
		`From`:         m.from,
		`To`:           m.to[0],
		`Subject`:      m.subject,
		`Content-Type`: `text/html; charset=UTF-8`,
	}

	builder := new(strings.Builder)

	for k, v := range header {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	builder.WriteString("\r\n" + m.content)

	m.msg = []byte(builder.String())
}

func (m *Mail) Send() error {
	return m.sendMailUsingTLS()

}

// TLS 发送邮件
func (m *Mail) sendMailUsingTLS() (err error) {

	defer func() {
		err = m.mailClient.Quit()
	}()

	m.buildMsg()

	if m.auth != nil {
		if ok, _ := m.mailClient.Extension("AUTH"); ok {
			if err = m.mailClient.Auth(m.auth); err != nil {
				return err
			}
		}
	}

	if err = m.mailClient.Mail(m.mailUser); err != nil {
		return err
	}

	for _, addr := range m.to {
		if err = m.mailClient.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := m.mailClient.Data()

	if err != nil {
		return err
	}

	defer func() {
		err = w.Close()
	}()

	_, err = w.Write(m.msg)

	return

}
