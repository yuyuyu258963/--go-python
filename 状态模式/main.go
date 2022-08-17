package main

import "fmt"

/**
* 状态接口
 */
type State interface {
	doAction(c Context)
}

type StartState struct{}

func (s *StartState) doAction(c Context) {
	fmt.Println("StartState")
	c.setState(s)
}

type StopState struct{}

func (s *StopState) doAction(c Context) {
	fmt.Println("StopState")
	c.setState(s)
}

/**
 * Context类
 */
type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) getState() *State {
	return &c.state
}

func statePatternDemo() {
	context := &Context{}
	startState := &StartState{}
	stopState := &StopState{}
	startState.doAction(*context)
	stopState.doAction(*context)
}

func main() {
	statePatternDemo()
}
