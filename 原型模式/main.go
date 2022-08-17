package main

import "fmt"

/**
 * 抽象类接口
 */
type Shape struct {
	Id   string
	Type string
}

func (s *Shape) getType() string {
	return s.Type
}

func (s *Shape) getId() string {
	return s.Id
}

type ShapeT interface {
	Draw()
	Clone() ShapeT
}

type Rectangle struct {
	Shape
}

type Square struct {
	Shape
}

func (s *Square) Draw() {
	fmt.Printf("Drawing a square with id %s\n", s.getId())
}

/**
 * 深拷贝，核心
 */
func (s *Square) Clone() ShapeT {
	return &Square{Shape{
		Id:   s.Id,
		Type: s.Type,
	}}
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle with id %s\n", r.getId())
}

/**
 * 深拷贝，核心
 */
func (r *Rectangle) Clone() ShapeT {
	return &Rectangle{Shape{
		Id:   r.Id,
		Type: r.Type,
	}}
}

type ShapeCache struct {
	shapes map[string]ShapeT
}

func (s *ShapeCache) loadCache() {
	s.shapes = make(map[string]ShapeT)
	s.shapes["1"] = &Rectangle{
		Shape{
			Id:   "1",
			Type: "rectangle",
		},
	}
	s.shapes["2"] = &Square{
		Shape{
			Id:   "2",
			Type: "square",
		},
	}
}

func (s *ShapeCache) getShape(id string) ShapeT {
	return s.shapes[id].Clone()
}

func test() {
	shapeCache := ShapeCache{}
	shapeCache.loadCache()
	rect := shapeCache.getShape("1").(*Rectangle)
	fmt.Printf("Shape %s\n", rect.getType())
	square := shapeCache.getShape("2").(*Square)
	fmt.Printf("Shape %s\n", square.getType())
}

func main() {
	test()
}
