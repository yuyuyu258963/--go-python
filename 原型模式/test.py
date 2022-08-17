class Shape(object):
  def __init__(self, Id, type):
    self.id = Id
    self.type = type
  def getId(self):
    return self.id
  def getType(self):
    return self.type
  def clone(self):
    """
      ## 原型模式的关键在于clone方法
    """
    return super(Shape, self).clone()

class Circle(Shape):
  def __init__(self, Id, type, radius):
    super(Circle, self).__init__(Id, type)
    self.radius = radius
  def draw(self):
    print("Circle with radius %d" % self.radius)

class Square(Shape):
  def __init__(self, Id, type, side):
    super(Square, self).__init__(Id, type)
    self.side = side
  def draw(self):
    print("Square with side %d" % self.side)

class ShapeCache(object):
  def __init__(self):
    self.shapeMap = {}
  def getShape(self, Id):
    shape = self.shapeMap.get(Id, None)
    if shape == None:
      raise Exception("Shape not found")
    return shape
  def loadCache(self):
    circle = Circle("1", "Circle", 3)
    square = Square("2", "Square", 4)
    self.shapeMap["1"] = circle
    self.shapeMap["2"] = square

def PrototypePatternDemo():
  shapeCache = ShapeCache()
  shapeCache.loadCache()
  circle = shapeCache.getShape("1")
  circle.draw()
  square = shapeCache.getShape("2")
  square.draw()

if __name__ == '__main__':
  PrototypePatternDemo()