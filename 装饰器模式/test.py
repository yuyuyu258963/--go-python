
class Shape(object):
  """
    ## 形状抽象类
  """
  def __init__(self, shapeName):
    self.shapeName = shapeName
  def draw(self,):
    print(f"Draw {self.shapeName}")

class Circle(Shape):
  def __init__(self, shapeName, radius):
    super().__init__(shapeName)
    self.radius = radius

class Rectangle(Shape):
  def __init__(self, shapeName, width, height):
    super().__init__(shapeName)
    self.width = width
    self.height = height

class ShapeDecorator(Shape):
  """
    ## 装饰器抽象类
  """
  def __init__(self, shape):
    self.shape = shape
  def draw(self,):
    self.shape.draw()

class ColorDecorator(ShapeDecorator):
  """
    ## 颜色装饰器
  """
  def __init__(self, shape, borderColor="red"):
    super().__init__(shape)
    self.borderColor = borderColor
  def draw(self):
    self.setRedBorder()
    self.shape.draw()
  def setRedBorder(self):
    print(f"Border Color: {self.borderColor}")

def test():
  c = Circle("Circle", 10)
  redCircle = ColorDecorator(c, "red")
  redCircle.draw()
  r = Rectangle("Rectangle", 10, 10)
  blueRectangle = ColorDecorator(r, "blue")
  blueRectangle.draw()

if __name__ == "__main__":
  test()

