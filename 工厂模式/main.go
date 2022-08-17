package main

import "fmt"

type ShapeDrawer interface {
	Draw()
}

type shapeCommFuncs struct {
	Name string
}

type Circle struct {
	x, y, radius int
	shapeCommFuncs
}

func (t *shapeCommFuncs) Draw() {
	fmt.Printf("绘制%s\n", t.Name)
}

type Square struct {
	x, y, length int
	shapeCommFuncs
}

type Rectangle struct {
	x, y, width, height int
	shapeCommFuncs
}

/**
 * 工厂模式创建对象
 */
func createFactory(shapeName string) ShapeDrawer {
	switch shapeName {
	case "Circle":
		return &Circle{
			shapeCommFuncs: shapeCommFuncs{Name: shapeName},
		}
	case "Square":
		return &Square{
			shapeCommFuncs: shapeCommFuncs{Name: shapeName},
		}
	case "Rectangle":
		return &Rectangle{
			shapeCommFuncs: shapeCommFuncs{Name: shapeName},
		}
	default:
		return nil
	}
}

func test() {
	c := createFactory("Circle")
	s := createFactory("Square")
	c.Draw()
	s.Draw()
}

func main() {
	test()
}
