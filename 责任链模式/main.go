package main

import "fmt"

var status = 0

type HandlerFunc func()

type HandlersChain []HandlerFunc

/**
* 路由组件
 */
type RouteGroup struct {
	Handlers HandlersChain
	index    int8
}

/**
* 添加中间件，将其组成链式
 */
func (group *RouteGroup) Use(handlers ...HandlerFunc) {
	group.Handlers = append(group.Handlers, handlers...)
}

/**
* 链式执行
 */
func (group *RouteGroup) Next() {
	for group.index < int8(len(group.Handlers)) {
		group.Handlers[group.index]()
		if status == 1 {
			fmt.Printf("执行失败，请重试\n")
			break
		}
		group.index++
	}
	group.index = 0
}

/**
* 中间件1
 */
func middleWare1() {
	fmt.Printf("middleWare1\n")
}

/**
* 中间件1
 */
func middleWare2() {
	fmt.Printf("middleWare2\n")
	status = 1
}

func test() {
	r := &RouteGroup{}
	r.Use(middleWare1, middleWare2)
	r.Next()

}

func main() {
	test()
}
