package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) {
	fmt.Println("Paying with PayPal:", amount)
}

type Stripe struct{}

func (s *Stripe) Pay(amount float64) {
	fmt.Println("Paying with Stripe:", amount)
}

type PaymentContext struct {
	Strategy PaymentStrategy
}

func (pc *PaymentContext) Execute(amount float64) {
	pc.Strategy.Pay(amount)
}

// 通过策略模式，可在运行时灵活切换算法行为。
// 应用场景：如支付网关的切换（PayPal、Stripe 等）。
func main() {
	ctx := PaymentContext{
		Strategy: &PayPal{},
	}
	ctx.Execute(250.0)
}
