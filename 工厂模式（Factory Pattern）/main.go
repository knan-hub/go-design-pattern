package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending email:", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

func GetNotifier(channel string) Notifier {
	switch channel {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SMSNotifier{}
	default:
		return nil
	}
}

// 工厂模式为对象的创建提供一种高度抽象且简洁的方式。
// 应用场景：当系统中存在多种通知渠道（如电子邮件、短信、推送等）时，通过工厂模式可统一接口便捷创建各种通知对象。
func main() {
	notifier := GetNotifier("email")
	notifier.Send("Hello, world!")
}
