package main

import "fmt"

type ReadFile interface {
	Read(filePath string)
	Accept(visitor Visitor)
}

type ReadPdfFile struct{}

func (r *ReadPdfFile) Read(filePath string) {
	fmt.Printf("读取pdf文件：%s\n", filePath)
}

/**
* 接收访问者
 */
func (r *ReadPdfFile) Accept(visitor Visitor) {
	visitor.VistorPdfFile(r)
}

type ReadTxtFile struct{}

func (r *ReadTxtFile) Read(filePath string) {
	fmt.Printf("读取txt文件：%s\n", filePath)
}

func (r *ReadTxtFile) Accept(visitor Visitor) {
	visitor.VistorTxtFile(r)
}

type Visitor interface {
	VistorPdfFile(r *ReadPdfFile)
	VistorTxtFile(r *ReadTxtFile)
}

/**
 * 读取文件类
 */
type ExactFile struct{}

func (e *ExactFile) VistorPdfFile(r *ReadPdfFile) {
	fmt.Println("读取pdf文件")
}

func (e *ExactFile) VistorTxtFile(r *ReadTxtFile) {
	fmt.Println("读取txt文件")
}

/**
* 压缩文件类
 */
type CompressionFile struct{}

func (c *CompressionFile) VistorPdfFile(r *ReadPdfFile) {
	fmt.Println("压缩pdf文件")
}

func (c *CompressionFile) VistorTxtFile(r *ReadTxtFile) {
	fmt.Println("压缩txt文件")
}

func test() {
	fileList := []ReadFile{
		&ReadPdfFile{},
		&ReadTxtFile{},
		&ReadTxtFile{},
	}
	extract := ExactFile{}
	for _, f := range fileList {
		f.Accept(&extract)
	}
	fmt.Println("------------------------------------")
	compress := CompressionFile{}
	for _, f := range fileList {
		f.Accept(&compress)
	}
}

func main() {
	test()
}
