class Talking:
  """
    ## 父类
  """
  def __init__(self):
    pass
  def Answer(self):
    pass
  def Answer2(self):
    pass
  def say(self):
    print("Hello World")
    self.Answer()
    print("hi j")
    self.Answer2()

class Person(Talking):
  def Answer(self):
    print("hi ")
  def Answer2(self):
    print("what are you saying")

def test():
  person = Person()
  person.say()

if __name__ == '__main__':
  test()