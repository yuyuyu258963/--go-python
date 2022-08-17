class RouteGroup(object):
  def __init__(self):
    self.handlers = []
    self.index = 0
  def Use(self, *args):
    self.handlers += [*args]
  def Next(self):
    index = self.index
    for task in self.handlers[index:]:
      task()
      self.index += 1

def middleware1():
  print('middleware1')

def middleware2():
  print('middleware2')

def test():
  routeGroup = RouteGroup()
  routeGroup.Use(middleware1, middleware2)
  routeGroup.Next()

if __name__ == '__main__':
  test()  

