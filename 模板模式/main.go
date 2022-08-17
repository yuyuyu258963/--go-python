package main

import "fmt"

/**
* 问答
 */
type Talking struct {
	Answer  func()
	Answer2 func()
}

func (t *Talking) say() {
	fmt.Printf("Hello World\n")
	t.Answer()
	fmt.Printf("hi j\n")
	t.Answer2()
}

/**
* 相当于是子类
 */
type Person struct {
	Talking
}

func (p *Person) Answer() {
	fmt.Printf("hi\n")
}
func (p *Person) Answer2() {
	fmt.Printf("what are you saying\n")
}
func test() {
	talking := &Person{}
	talking.Talking.Answer = talking.Answer
	talking.Talking.Answer2 = talking.Answer2
	talking.say()
}

func main() {
	test()
}
