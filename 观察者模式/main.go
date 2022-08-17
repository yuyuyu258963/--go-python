package main

import "fmt"

type PurchaseOperFunc func(status string, data string) (res bool, err error)

var PurchaseOperFuncArr = []PurchaseOperFunc{
	create,
	isDeleted,
	apply,
}

/**
* 用于创建的观察者
 */
func create(status string, data string) (res bool, err error) {
	if status == "create" {
		fmt.Printf("create %s\n", data)
		return true, nil
	}
	return false, nil
}

/**
* 用于删除的观察者
 */
func isDeleted(status string, data string) (res bool, err error) {
	if status == "delete" {
		fmt.Printf("delete %s\n", data)
		return true, nil
	}
	return false, nil
}

/**
* 履行的观察者
 */
func apply(status string, data string) (res bool, err error) {
	if status == "apply" {
		fmt.Printf("apply %s\n", data)
		return true, nil
	}
	return false, nil
}

func test() {
	status := "create"
	data := "test"
	for _, v := range PurchaseOperFuncArr {
		res, err := v(status, data)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		if res {
			fmt.Printf("success\n")
		}
	}
}

func main() {
	test()
}
