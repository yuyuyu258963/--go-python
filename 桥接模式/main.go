package main

import "fmt"

/**
*桥接的抽象类
 */
type DrawAPI interface {
	drawCircle(x, y, radius int)
}

type RedCircle struct {
}

func (rc *RedCircle) drawCircle(x, y, radius int) {
	fmt.Printf("draw red circle [x:%d, y:%d, radius:%d]\n", x, y, radius)
}

type GreenCircle struct{}

func (rc *GreenCircle) drawCircle(x, y, radius int) {
	fmt.Printf("draw green circle [x:%d, y:%d, radius:%d]\n", x, y, radius)
}

// 形状的抽象类
type Shape struct {
	DrawAPI
}

func (s *Shape) shape(drawAPI DrawAPI) {
	s.DrawAPI = drawAPI
}

type Circle struct {
	x, y, radius int
	Shape
}

func (c *Circle) draw() {
	c.drawCircle(c.x, c.y, c.radius)
}

func (c *Circle) initShape(x, y, radius int, drawAPI DrawAPI) {
	c.x = x
	c.y = y
	c.radius = radius
	c.shape(drawAPI)
}

func test() {
	c := &Circle{}
	c.initShape(1, 2, 3, &GreenCircle{})
	c.draw()
}

func main() {
	test()
}
