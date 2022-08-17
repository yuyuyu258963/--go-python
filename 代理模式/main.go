package main

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (r *RealImage) LoadImage() {
	fmt.Printf("Loading image: %s\n", r.fileName)
}

func (rImage *RealImage) Display() {
	fmt.Printf("Displaying %s\n", rImage.fileName)
}

func initRealImage(fileName string) *RealImage {
	realImage := &RealImage{fileName}
	realImage.LoadImage()
	return realImage
}

type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func (pImage *ProxyImage) Display() {
	if pImage.realImage == nil {
		pImage.realImage = &RealImage{fileName: pImage.fileName}
	}
	pImage.realImage.Display()
}

func test() {
	image := initRealImage("test.jpg")
	// 要从磁盘加载
	image.Display()
	// 不需要从磁盘加载
	image.Display()
}

func main() {
	test()
}
