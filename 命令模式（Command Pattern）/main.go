package main

import "fmt"

type Command interface {
	Execute()
}

type PrintCommand struct {
	Msg string
}

func (p *PrintCommand) Execute() {
	fmt.Println(p.Msg)
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) Run() {
	for _, c := range i.commands {
		c.Execute()
	}
}

// 命令模式将请求封装为对象，使系统支持请求排队、操作日志和参数化处理等特性。
// 应用场景：作业调度或队列系统。
func main() {
	invoker := &Invoker{}
	invoker.AddCommand(&PrintCommand{Msg: "Hello, World!"})
	invoker.AddCommand(&PrintCommand{Msg: "Hello, Go!"})
	invoker.Run()
}
