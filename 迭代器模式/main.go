package main

import (
	"errors"
)

/**
* 迭代器接口
 */
type Iterator interface {
	hasNext() bool
	next() (string, error)
}

type Container interface {
	getIterator() Iterator
}

// 内部实现迭代器
type NameRepository struct {
	Names []string
}

func (r *NameRepository) getIterator() Iterator {
	return &NameIterator{
		Index: -1,
		Names: r.Names,
	}
}

type NameIterator struct {
	Index int
	Names []string
}

func (n *NameIterator) hasNext() bool {
	return n.Index < len(n.Names)-1
}

func (n *NameIterator) next() (string, error) {
	if n.hasNext() {
		n.Index++
		return n.Names[n.Index], nil
	}
	return "", errors.New("no next")
}

func IteratorPatternDemo() {
	names := []string{"yyy", "cy", "ly", "cc"}
	nameRepository := &NameRepository{Names: names}
	iterator := nameRepository.getIterator()
	for {
		if name, err := iterator.next(); err == nil {
			println(name)
		} else {
			break
		}
	}
}

func main() {
	IteratorPatternDemo()
}
