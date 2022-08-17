package main

type Burger struct {
}

func (c *Burger) Pack() string {
	packing := Wrapper{}
	return packing.Pack()
}

type VegBurger struct {
	Burger
}

func (v *VegBurger) Price() float64 {
	return float64(3.8)
}
func (v *VegBurger) Name() string {
	return "VegBurger"
}

type ChickenBurger struct {
	Burger
}

func (c *ChickenBurger) Price() float64 {
	return float64(2.8)
}

func (c *ChickenBurger) Name() string {
	return "ChickenBurger"
}

type ColdDrink struct {
}

func (c *ColdDrink) Pack() string {
	packing := Bottle{}
	return packing.Pack()
}

type Pepsi struct {
	ColdDrink
}

func (c *Pepsi) Price() float64 {
	return float64(1.8)
}

func (c *Pepsi) Name() string {
	return "Pepsi"
}

type Coke struct {
	ColdDrink
}

func (c *Coke) Price() float64 {
	return float64(0.8)
}

func (c *Coke) Name() string {
	return "Coke"
}
