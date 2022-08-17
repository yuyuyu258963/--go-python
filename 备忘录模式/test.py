class Memento:
  """
    ## 要记录的状态
  """
  def __init__(self, state):
    self.state = state
  def GetState(self):
    return self.state

class Originator:
  """
    ## 行为对象，可加载存储的状态
  """
  def __init__(self):
    self.state = None
  def SetState(self, state):
    self.state = state
  def GetState(self):
    return self.state
  def saveStateToMemento(self):
    return Memento(self.state)
  def loadStateFromMemento(self, mem):
    self.state = mem.GetState()

class CareTaker:
  """
    ## 用于记录状态
  """
  def __init__(self):
    self.mementoList = []
  def addMemento(self, memento):
    self.mementoList.append(memento)
  def getMemento(self, index):
    if index > len(self.mementoList):
      raise Exception("index out of range")
    return self.mementoList[index]

def test():
  original = Originator()
  caretaker = CareTaker()
  original.SetState("State 1")
  caretaker.addMemento(original.saveStateToMemento())
  original.SetState("State 2")
  caretaker.addMemento(original.saveStateToMemento())
  original.SetState("State 3")
  
  print(f"current state: {original.GetState()}")
  original.loadStateFromMemento(caretaker.getMemento(0))
  print(f"first state: {original.GetState()}")
  original.loadStateFromMemento(caretaker.getMemento(1))
  print(f"second state: {original.GetState()}")

if __name__ == '__main__':
  test()