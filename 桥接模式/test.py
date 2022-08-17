class DrawApI:
  def drawCircle(self, x, y, radius):
    print("Drawing circle at ({},{}) with radius {}".format(x, y, radius))

class RedCircle(DrawApI):
  def drawCircle(self, x, y, radius):
    print("Drawing red circle at ({},{}) with radius {}".format(x, y, radius))

class BlueCircle(DrawApI):
  def drawCircle(self, x, y, radius):
    print("Drawing blue circle at ({},{}) with radius {}".format(x, y, radius))

class Shape():
  def __init__(self, drawAPI):
    self.drawAPI = drawAPI
  def shape(self, drawAPI):
    self.drawAPI = drawAPI
  def draw():
    pass

class Circle(Shape):
  def __init__(self, x, y, radius, drawAPI):
    super().__init__(drawAPI)
    self.x = x
    self.y = y
    self.radius = radius
  def draw(self):
    self.drawAPI.drawCircle(self.x, self.y, self.radius)

def test():
  c = Circle(1, 2, 3, RedCircle())
  c.draw()

if __name__ == '__main__':
  test()