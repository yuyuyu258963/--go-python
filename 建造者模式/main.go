package main

import "fmt"

/**
* 建造者
 */
type MealBuilder struct{}

func (builder *MealBuilder) PrepareVegMeal() Meal {
	meal := Meal{}
	meal.addItem(&VegBurger{})
	meal.addItem(&Coke{})
	return meal
}

func (builder *MealBuilder) prepareNonVegMeal() Meal {
	meal := Meal{}
	meal.addItem(&ChickenBurger{})
	meal.addItem(&Pepsi{})
	return meal
}

func test() {
	builder := MealBuilder{}
	vegMeal := builder.PrepareVegMeal()
	NoeVegMeal := builder.prepareNonVegMeal()
	vegMeal.showItems()
	fmt.Printf("Cost of the veg meal is: %.2f\n", vegMeal.getCost())
	NoeVegMeal.showItems()
	fmt.Printf("Cost of the Non-Veg meal is: %.2f\n", NoeVegMeal.getCost())
}

func main() {
	test()
}
