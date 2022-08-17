class Shape:
  """
    ## Shape 接口
  """
  def draw(self):
    pass

class Circle(Shape):
  def draw(self):
    print("Circle.draw")

class Square(Shape):
  def draw(self):
    print("Square.draw")

class ShapeMaker:
  """
    ## 外观类
  """
  def __init__(self):
    self.circle = Circle()
    self.square = Square()
  def drawCircle(self):
    self.circle.draw()
  def drawSquare(self):
    self.square.draw()

def test():
  shapeMaker = ShapeMaker()
  shapeMaker.drawCircle()
  shapeMaker.drawSquare()

if __name__ == '__main__':
  test()