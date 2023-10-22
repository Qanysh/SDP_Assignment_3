package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Data int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: 42}
	})
	return instance
}

func main() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	singleton1.Data = 10

	fmt.Printf("Данные singleton1: %d\n", singleton1.Data)
	fmt.Printf("Данные singleton2: %d\n", singleton2.Data)

	if singleton1 == singleton2 {
		fmt.Println("singleton1 и singleton2 - это один и тот же экземпляр Singleton.")
	} else {
		fmt.Println("singleton1 и singleton2 - это разные экземпляры Singleton.")
	}
}
