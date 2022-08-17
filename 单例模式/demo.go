package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	Counter int
}

var once sync.Once
var instance *singleton

// 获得单例对象
func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
func testSingleton() {
	ywh := getInstance()
	yhm := getInstance()
	ywh.Counter++
	fmt.Println(ywh)
	fmt.Println(yhm)
}

func main() {
	testSingleton()
}
