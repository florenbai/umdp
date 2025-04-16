package channel

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
	"gopkg.in/gomail.v2"
	"strconv"
)

type Email struct {
	*gomail.Dialer
}

type EmailOptions struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	SmtpServer string `json:"smtpServer"`
	SmtpPort   string `json:"smtpPort"`
}

type EmailParam struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type EOptions func(email *EmailOptions)

func NewEmailChannel() MessageChannel {
	return Email{}
}

func (email Email) SetConfig(config string) (MessageChannel, error) {
	var options EmailOptions
	json.Unmarshal([]byte(config), &options)
	port, _ := strconv.Atoi(options.SmtpPort)
	return &Email{
		gomail.NewDialer(options.SmtpServer, port, options.Username, options.Password),
	}, nil
}

func (email Email) Handle(ctx context.Context, parameters Parameters, tp TemplateParameters, retry int64) error {
	if len(parameters.Receiver) < 1 {
		return errors.New("必须含有一个接收人")
	}
	var emailParam EmailParam
	err := json.Unmarshal([]byte(tp.Config), &emailParam)
	if err != nil {
		return errors.New("邮箱格式不正确")
	}
	err = validator.New().Struct(emailParam)
	if err != nil {
		return errors.New("邮箱格式不正确")
	}
	if gjson.Valid(parameters.Variables) {
		emailParam.Title = ParameterMatch(parameters.Variables, emailParam.Title)
		emailParam.Content = ParameterMatch(parameters.Variables, emailParam.Content)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", email.Username)
	m.SetHeader("To", parameters.Receiver...)
	if len(parameters.Cc) > 0 {
		m.SetHeader("Cc", parameters.Cc...)
	}

	m.SetHeader("Subject", emailParam.Title)
	//发送html格式邮件。
	m.SetBody("text/html", emailParam.Content)
	return email.DialAndSend(m)
}
