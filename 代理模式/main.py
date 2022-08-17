class Image(object):
  """
    ## 图像的抽象类
  """
  def __init__(self, fileName):
    self.fileName = fileName
    pass
  def Display(self):
    pass

class RealImage(Image):
  def __init__(self, fileName):
    super().__init__(fileName)
    self.LoadImage()
  def LoadImage(self):
    print(f"Loading {self.fileName}")
  def Display(self):
    print(f"Displaying {self.fileName}")

class ProxyImage(Image):
  """
    ## 图像的代理类
  """
  def __init__(self, fileName):
    super().__init__(fileName)
    self.realImage = None
  def Display(self):
    if self.realImage is None:
      self.realImage = RealImage(self.fileName)
    self.realImage.Display()

def test():
  image = ProxyImage("test.jpg")
  image.Display()
  image.Display()

if __name__ == '__main__':
  test()