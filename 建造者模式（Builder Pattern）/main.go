package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: &User{}}
}

func (ub *UserBuilder) SetName(name string) *UserBuilder {
	ub.user.Name = name
	return ub
}

func (ub *UserBuilder) SetEmail(email string) *UserBuilder {
	ub.user.Email = email
	return ub
}

func (ub *UserBuilder) SetAge(age int) *UserBuilder {
	ub.user.Age = age
	return ub
}

func (ub *UserBuilder) Build() *User {
	return ub.user
}

// 当对象构造过程复杂，或有大量可选参数时，建造者模式极大提升易用性与清晰度。
// 应用场景：构建复杂的配置对象或 HTTP 请求。
func main() {
	user := NewUserBuilder().SetName("knan").SetEmail("knan@mail.com").SetAge(18).Build()
	fmt.Println(user)
}
