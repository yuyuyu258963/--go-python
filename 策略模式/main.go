package main

import "fmt"

const (
	Common = "Common"
	Win    = "Win"
)

/**
* 获取商品类型
 */
func getProductType(hscode string) string {
	if hscode == "11" {
		return Common
	} else {
		return Win
	}
}

/**
* 税费计算函数
 */
type TaxComputeFunc func(price int64, qty int64) (taxPrice int64)

var TaxComputeFunMap = map[int]map[string]TaxComputeFunc{
	0: map[string]TaxComputeFunc{
		Common: common,
		Win:    win,
	},
	1: map[string]TaxComputeFunc{
		Common: common,
		Win:    win,
	},
}

/**
* 计算普通商品税费
 */
func common(price int64, qty int64) (taxPrice int64) {
	radio := 0.1
	fmt.Printf("计算普通商品税费")
	taxPrice = int64(float64(price*qty) * radio)
	return taxPrice
}

/**
* 计算窗口商品税费
 */
func win(price int64, qty int64) (taxPrice int64) {
	radio := 0.2
	fmt.Printf("计算窗口商品税费")
	taxPrice = int64(float64(price*qty) * radio)
	return taxPrice
}

/**
*计算税费
*@param price 商品价格
 */
func ComputeTaxPrice(withTax int, productType string, price int64, qty int64) {
	if taxFunc, ok := TaxComputeFunMap[withTax][productType]; ok {
		taxPrice := taxFunc(price, qty)
		fmt.Printf("税费为:%d", taxPrice)
	} else {
		fmt.Printf("没有找到税费计算函数")
	}
}

func test() {
	withTax := 0
	var price, qty int64 = 1000, 3
	productType := getProductType("11")
	ComputeTaxPrice(withTax, productType, price, qty)
}

func main() {
	test()
}
