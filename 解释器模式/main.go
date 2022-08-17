package main

import "fmt"

/**
* 内容信息
 */
type Context struct {
	action  string
	context string
}

/**
* 翻译接口
 */
type InterPreter interface {
	Interpret(context Context)
}

type MusicInterpreter struct {
}

/**
* 音乐翻译器
 */
func (m *MusicInterpreter) Interpret(context Context) {
	fmt.Printf("%s 中的 %s 的意思是温婉尔雅\n", context.action, context.context)
}

/**
* 舞蹈翻译器
 */
type DanceInterpreter struct {
}

func (d *DanceInterpreter) Interpret(context Context) {
	fmt.Printf("%s 中的 %s 意境很ok \n", context.action, context.context)
}

func test() {
	actionList := []Context{
		{action: "music", context: "高音"},
		{action: "music", context: "低音"},
		{action: "dance", context: "跳跃"},
		{action: "dance", context: "旋转"},
	}
	for _, c := range actionList {
		if c.action == "music" {
			(&MusicInterpreter{}).Interpret(c)
		} else if c.action == "dance" {
			(&DanceInterpreter{}).Interpret(c)
		}
	}
}

func main() {
	test()
}
