package main

import "fmt"

/**
* Shape的接口
 */
type Shape interface {
	draw()
}

type Circle struct{}

func (c *Circle) draw() {
	fmt.Println("draw Circle")
}

type Rectangle struct{}

func (r *Rectangle) draw() {
	fmt.Println("draw Rectangle")
}

type Square struct{}

func (s *Square) draw() {
	fmt.Println("draw Square")
}

/**
* 外观类
 */
type ShaperMaker struct {
	circle Circle
	rect   Rectangle
	square Square
}

func (m *ShaperMaker) drawCircle() {
	m.circle.draw()
}

func (m *ShaperMaker) drawSquare() {
	m.square.draw()
}

func (m *ShaperMaker) drawRectangle() {
	m.rect.draw()
}

func initShaperMaker() *ShaperMaker {
	return &ShaperMaker{
		circle: Circle{},
		rect:   Rectangle{},
		square: Square{},
	}
}

func test() {
	shapeMarker := initShaperMaker()
	shapeMarker.drawCircle()
	shapeMarker.drawSquare()
	shapeMarker.drawRectangle()
}

func main() {
	test()
}
