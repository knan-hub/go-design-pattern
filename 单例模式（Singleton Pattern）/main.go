package main

import (
	"fmt"
	"sync"
)

type Config struct {
	DatabaseURL string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			DatabaseURL: "localhost:3306",
		}
	})
	return instance
}

// 单例模式确保某一类型仅有一个实例，并提供统一访问入口。
func main() {
	config1 := GetConfig()
	config2 := GetConfig()

	fmt.Println(config1 == config2)
}
