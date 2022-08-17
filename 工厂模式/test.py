class Shape():
  """
    ## 形状的抽象类
  """
  def __init__(self, shapeName):
    self.shapeName = shapeName
  def draw(self):
    print("drawing " + self.shapeName)

class Circle(Shape):
  def __init__(self, shapeName, radius):
    super().__init__(shapeName)
    self.radius = radius

class Square(Shape):
  def __init__(self, shapeName, xWidth):
    super().__init__(shapeName)
    self.xWidth = xWidth

class Rectangle(Shape):
  def __init__(self, shapeName, xWidth, yHeight):
    super().__init__(shapeName)
    self.xWidth = xWidth
    self.yHeight = yHeight

class ShapeFactory():
  """
    ## 工厂模式实现
  """
  shapeConfig = {
    "circle": Circle,
    "square": Square,
    "rectangle": Rectangle
  }
  @classmethod
  def getShape(self, shapeName, *args, **kwargs):
    obj = self.shapeConfig.get(shapeName, None)
    if obj is None:
      raise Exception("Shape not found")
    return obj(shapeName, *args, **kwargs)

def test():
  circle = ShapeFactory().getShape("circle", 10)
  square = ShapeFactory().getShape("square", 10)
  circle.draw()
  square.draw()

if __name__ == '__main__':
  test()
