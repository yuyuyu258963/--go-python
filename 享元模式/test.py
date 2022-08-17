import random

class Shape():
  """
    ## 抽象类
  """
  def draw(self):
    print("draw")

class Circle(Shape):
  """
    ## 实现接口的实体类
  """
  def __init__(self, x, y, radius, color):
    self.x = x
    self.y = y
    self.radius = radius
    self.color = color
  def draw(self):
    print(f"draw circle at {self.x},{self.y} with radius {self.radius} and color {self.color}")

class ShapeFactory():
  def __init__(self):
    self.circleMap = dict()
  def getCircles(self, name):
    if name not in self.circleMap:
      print(f"create {name} circle ")
      self.circleMap[name] = Circle(random.randint(0, 100), random.randint(0, 100), random.randint(0, 100), name)
    return self.circleMap[name]

def FlyWeightPatternDemo():
  factory = ShapeFactory()
  colors = ["red", "green", "blue", "yellow", "pink", "black", "white"]
  for _ in range(5):
    circle = factory.getCircles(colors[random.randint(0, len(colors)-1)])
    circle.draw()

if __name__ == '__main__':
  FlyWeightPatternDemo()
