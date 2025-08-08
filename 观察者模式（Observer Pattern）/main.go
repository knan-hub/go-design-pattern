package main

import "fmt"

type Observer interface {
	Update(data string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(data string) {
	for _, o := range s.observers {
		o.Update(data)
	}
}

type Logger struct{}

func (l *Logger) Update(data string) {
	fmt.Println("Logging:", data)
}

// 当一个对象状态变化时，观察者模式可通知并更新所有相关对象。
// 应用场景：实时系统，如 WebSocket 广播、通知分发等。
func main() {
	subject := &Subject{}
	logger := &Logger{}
	subject.Register(logger)
	subject.Notify("Hello, world!")
}
