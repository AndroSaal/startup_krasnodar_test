package internal

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/startup_krasnodar_test/src/pkg/config"
)

func (a *Auth) VerifyEmail(id int, code string) (bool, error) {
	fi := "internal.Auth.VerifyEmail"

	isVerified, err := a.RepositoryHandler.GetCodeFromEmail(id, code)

	if err != nil {
		a.logger.Debug(fmt.Sprintf("%s: %s", fi, err.Error()))
		return false, err
	}

	return isVerified, nil
}

// почта с которой будем отправлять писаьма с просьбой подтвердить email
type Mail struct {
	Config config.ServerMailAuthConf
	Logger *slog.Logger
}

func NewMailSender(config config.ServerMailAuthConf, log *slog.Logger) *Mail {
	return &Mail{
		Config: config,
		Logger: log,
	}
}

func (m *Mail) SendMail(toEmail, mailBody string) error {
	fi := "internal.Mail.SendMail"

	//созздаем клиента для отправки письма
	client, err := makeConnection(m, toEmail)
	if err != nil {
		m.Logger.Debug(fmt.Sprintf("%s: %s", fi, err.Error()))
		return err
	}
	//закрываем клиента
	defer client.Quit()

	//создаем writerа
	writer, err := client.Data()
	if err != nil {
		m.Logger.Debug(fmt.Sprintf("%s: %s", fi, err.Error()))
		return err
	}
	//закрываем writer
	defer writer.Close()

	//отправка письма
	writer.Write([]byte(mailBody))

	return nil

}

func makeConnection(m *Mail, toEmail string) (*smtp.Client, error) {
	fi := "internal.makeConnection"
	//аутенстификация серверной почты
	auth := smtp.PlainAuth("", m.Config.Login, m.Config.Password, m.Config.Host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.Config.Host,
	}

	//создаем соединение с нужным smtp сервером
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", m.Config.Host, m.Config.Port), tlsConfig)
	if err != nil {
		return nil, err
	}

	//создание smtp клиента
	client, err := smtp.NewClient(conn, m.Config.Host)
	if err != nil {

		return nil, err
	}

	//аторизируем клиента
	if err := client.Auth(auth); err != nil {
		return nil, err
	}

	// **FROM**
	if err := client.Mail(m.Config.Login); err != nil {
		return nil, err
	}

	// 	**TO**
	if err := client.Rcpt(toEmail); err != nil {
		return nil, err
	}

	//для трейса
	defer func(error) {
		if err != nil {
			m.Logger.Error(fi + ":" + err.Error())
		}
	}(err)

	return client, nil
}
