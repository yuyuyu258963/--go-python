package main

import "fmt"

type Item interface {
	Price() float64
	Name() string
	Pack() string
}

type Meal struct {
	Items []Item
}

func (m *Meal) addItem(item Item) {
	m.Items = append(m.Items, item)
}

func (m *Meal) getCost() (sum float64) {
	for _, item := range m.Items {
		sum += item.Price()
	}
	return sum
}

func (m *Meal) showItems() {
	fmt.Println("Items:")
	for _, item := range m.Items {
		fmt.Printf("Item: %s, Price: %.2f, Pack:%s\n", item.Name(), item.Price(), item.Pack())
	}
}
