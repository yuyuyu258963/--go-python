package main

import (
	"fmt"
	"time"
)

/**
* 中介类
 */
type ChatRoom struct{}

func (c ChatRoom) Send(user *User, message string) {
	fmt.Printf("%v %s say: %s\n", time.Now(), user.GetName(), message)
}

/**
* 用户类
 */
type User struct {
	Name string
}

func (u *User) GetName() string { return u.Name }

func (u *User) SetName(name string) { u.Name = name }

func (u *User) send(message string) {
	c := ChatRoom{}
	c.Send(u, message)
}

func MediatorPatternDemo() {
	rebort := &User{"Rebort"}
	john := &User{"John"}
	rebort.send("Hello, John")
	john.send("Hello, Rebot")
}

func main() {
	MediatorPatternDemo()
}
