package main

import "fmt"

/**
* 数据存储的对象
 */
type Memento struct {
	state string
}

func (m *Memento) GetState() string { return m.state }

type Originator struct {
	state string
}

func (m *Originator) GetState() string { return m.state }

func (m *Originator) SetState(state string) { m.state = state }

/**
* 保存状态
 */
func (m *Originator) saveStateToMemento() *Memento {
	return &Memento{m.state}
}

/**
* 加载状态
 */
func (m *Originator) getStateFromMemento(mem *Memento) {
	m.state = mem.GetState()
}

/**
* 备忘录模式核心，记录状态
 */
type CareTaker struct {
	mementoList []*Memento
}

func (c *CareTaker) add(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) get(index int) *Memento {
	return c.mementoList[index]
}

func test() {
	originator := &Originator{state: "State"}
	careTaker := &CareTaker{mementoList: make([]*Memento, 0)}
	originator.SetState("State == 1")
	careTaker.add(originator.saveStateToMemento())
	originator.SetState("State == 2")
	careTaker.add(originator.saveStateToMemento())

	fmt.Printf("current state: %v\n", originator.state)
	originator.getStateFromMemento(careTaker.get(0))
	fmt.Printf("current state: %v\n", originator.state)
	originator.getStateFromMemento(careTaker.get(1))
	fmt.Printf("current state: %v\n", originator.state)
}

func main() {
	test()
}
