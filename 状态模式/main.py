class State(object):
  """
    ## 状态的抽象类
  """
  def adAction(self, action):
    pass

class Context(object):
  def __init__(self):
    self.state = None
  def setState(self, state):
    self.state = state
  def getState(self):
    return self.state

class StartState(State):
  def adAction(self, action):
    print("Start state")
    action.setState(self )

class StopState(State):
  def adAction(self, action):
    print("Stop state")
    action.setState(self)

def test():
  context = Context()
  startState = StartState()
  stopState = StopState()
  startState.adAction(context)
  stopState.adAction(context)

if __name__ == '__main__':
  test()