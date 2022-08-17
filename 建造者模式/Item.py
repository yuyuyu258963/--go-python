class Item(object):
  def Price(self):
    return self.price
  def Name(self):
    return self.name
  def Pack(self):
    self.Pack()

class Meal():
  def __init__(self):
    self.items = []
  def addItem(self, item):
    self.items.append(item)
  def getCost(self):
    cost = 0.0
    for item in self.items:
      cost += item.Price()
    return cost
  def showItems(self):
    print(f"Items:")
    for item in self.items:
      print(f"Name::{item.Name()}, Price::{item.Price()}")