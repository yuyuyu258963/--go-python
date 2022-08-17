package main

import "fmt"

/**
* 命令接口
 */
type Command interface {
	Execute()
}

/**
* 移动命令
 */
type MoveCommand struct {
	x, y int
}

func (m *MoveCommand) Execute() {
	fmt.Printf("Move to %d, %d\n", m.x, m.y)
}

/**
* 攻击命令
 */
type AttackCommand struct {
	skill string
}

/**
* 攻击的动作执行
 */
func (a *AttackCommand) Execute() {
	fmt.Printf("Attack with %s\n", a.skill)
}

func AddCommand(action string) Command {
	if action == "attack" {
		return &AttackCommand{skill: "sword"}
	} else {
		return &MoveCommand{x: 1, y: 2}
	}
}

func test() {
	commandList := make([]Command, 0)
	commandList = append(commandList, AddCommand("attack"))
	commandList = append(commandList, AddCommand("move"))
	commandList = append(commandList, AddCommand("attack"))
	for _, c := range commandList {
		c.Execute()
	}
}

func main() {
	test()
}
