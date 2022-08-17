package main

import "fmt"

/**
* 抽象父类
 */
type Shape interface {
	Draw()
}

/**
* 共同的方法，属性
 */
type shapeCommFunc struct {
	name string
}

func (c *shapeCommFunc) Draw() {
	fmt.Printf("Draw： %s\n", c.name)
}

type Circle struct {
	x, y, radius int
	*shapeCommFunc
}

type Rectangle struct {
	x, y, width, height int
	*shapeCommFunc
}

/**
 * 装饰器
 */
type Decoder struct {
	shape Shape
	*ColorDecoder
}

/**
 * 颜色装饰器
 */
type ColorDecoder struct {
	color string
}

func (c *ColorDecoder) setBorderColor() {
	fmt.Printf("设置Border:%s\n", c.color)
}

func (c *Decoder) Draw() {
	c.setBorderColor()
	c.shape.Draw()
}

func test() {
	redCircle := &Decoder{&Circle{x: 1, y: 2, radius: 3, shapeCommFunc: &shapeCommFunc{"Circle"}}, &ColorDecoder{"red"}}
	redCircle.Draw()
	redRectangle := &Decoder{&Rectangle{x: 1, y: 2, width: 3, height: 4, shapeCommFunc: &shapeCommFunc{"Rectangle"}}, &ColorDecoder{"blue"}}
	redRectangle.Draw()
}

func main() {
	test()
}
