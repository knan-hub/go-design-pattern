package main

import (
	"fmt"
	"log"
)

type Service interface {
	Execute() string
}

type BaseService struct{}

func (b *BaseService) Execute() string {
	return "Base Service"
}

type LoggingDecorator struct {
	Wrapped Service
}

func (l *LoggingDecorator) Execute() string {
	log.Println("Before executing service")
	res := l.Wrapped.Execute()
	log.Println("After executing service")
	return res
}

// 装饰器模式允许在不改变原对象代码的前提下，动态为其添加新功能。
// 应用场景：为服务增加日志、监控、认证或重试逻辑。
func main() {
	svc := &LoggingDecorator{
		Wrapped: &BaseService{},
	}
	fmt.Println(svc.Execute())
}
