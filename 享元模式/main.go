package main

import (
	"fmt"
	"math/rand"
)

type Shape interface {
	Draw()
}

/**
 * 创建实现接口的实体类
 */
type Circle struct {
	x, y, radius int
	color        string
}

func (c *Circle) Draw() {
	fmt.Printf("Circle: x=%d, y=%d, radius=%d, color=%s\n", c.x, c.y, c.radius, c.color)
}

/**
 * 创建一个工厂，生成基于给定信息的实体类的对象
 */
type ShapeFactory struct {
	circleMap map[string]Shape
}

func (f *ShapeFactory) getCircle(name string) *Circle {
	if c, ok := f.circleMap[name]; ok {
		return c.(*Circle)
	} else {
		fmt.Printf("Creating circle of color : %s\n", name)
		circle := &Circle{color: name}
		f.circleMap[name] = circle
		return circle
	}
}

func FlyWeightPatternDemo() {
	colors := []string{"Red", "Green", "Blue", "White", "Black"}
	f := &ShapeFactory{circleMap: make(map[string]Shape)}
	for i := 0; i < 5; i++ {
		cIndex := rand.Intn(len(colors))
		circle := f.getCircle(colors[cIndex])
		circle.Draw()
	}
}

func main() {
	FlyWeightPatternDemo()
}
