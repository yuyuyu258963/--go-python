from pack import Wrapper, Bottle
from Item import Item

class Burger(Item):
  def Pack(self):
    return Wrapper().Pack() 

class VegBurger(Burger):
  def __init__(self):
    self.name = "Veg Burger"
    self.price = 1.5

class ChickenBurger(Burger):
  def __init__(self):
    self.name = "Chicken Burger"
    self.price = 2.5

class ClodDrink(Item):
  def Pack(self):
    return Bottle().Pack() 

class Coke(ClodDrink):
  def __init__(self):
    self.name = "Coke"
    self.price = 0.9

class Pepsi(ClodDrink):
  def __init__(self):
    self.name = "Pepsi"
    self.price = 0.4