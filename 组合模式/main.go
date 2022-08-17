package main

import "fmt"

const Separator = "--"

/**
 * @Description: 文件系统的抽象类,文件和文件夹都实现该接口
 */
type FileSystemNode interface {
	Display(Separator string)
}

/**
 * @Description: 文件的通用功能
 */
type FileCommonFunc struct {
	fileName string
}

/**
* @Description: 设置文件名
 */
func (f *FileCommonFunc) SetFileName(fileName string) {
	f.fileName = fileName
}

/**
 * @Description: 文件类
 */
type FileNode struct {
	FileCommonFunc
}

/**
* @Description: 显示文件内容
 */
func (f *FileNode) Display(Separator string) {
	fmt.Println(Separator + f.fileName + "   文件内容为: Hello world")
}

/**
 * @Description: 目录类
 */
type DirectoryNode struct {
	FileCommonFunc
	nodes []FileSystemNode
}

/**
* @Description: 显示目录中的所有内容
 */
func (d *DirectoryNode) Display(Separator string) {
	fmt.Println(Separator + d.fileName + "   文件夹内容为:")
	for _, node := range d.nodes {
		node.Display(Separator + Separator)
	}
}

/**
 * @Description: 添加目录或者文件
 */
func (d *DirectoryNode) Add(f FileSystemNode) {
	d.nodes = append(d.nodes, f)
}

func test() {
	root := &DirectoryNode{}
	root.SetFileName("root")
	var file1 = &FileNode{}
	file1.SetFileName("file1")
	var file2 = &FileNode{}
	file2.SetFileName("file2")
	root.Add(file1)
	root.Add(file2)
	root.Display(">>")
}

func main() {
	test()
}
