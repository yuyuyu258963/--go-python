class Observer(object):
  def __init__(self, task):
    task.add_observer(self)
  def update(self):
    pass

class Create(Observer):
  """
    ## 订单创建观察者
  """
  def __init__(self,task):
    super().__init__(task)
  def update(self, state):
    if state == 'create':
      print('create')

class isDelete(Observer):
  """
    ## 订单删除观察者
  """
  def __init__(self,task):
    super().__init__(task)
  def update(self, state):
    if state == 'delete':
      print('delete')

class apply(Observer):
  """
    ## 订单履行观察者
  """
  def __init__(self,task):
    super().__init__(task)
  def update(self, state):
    if state == 'apply':
      print('apply')

class Task(object):
  """
    ## 订单任务
  """
  def __init__(self, state):
    self.state = state
    self.observers = []
  def setState(self, state):
    self.state = state
    self.notify()
  def add_observer(self, observer):
    self.observers.append(observer)
  def notify(self):
    for observer in self.observers:
      observer.update(self.state)

def test():
  task = Task("create")
  Create(task)
  isDelete(task)
  apply(task)
  task.setState("apply")
  task.setState("delete")

if __name__ == '__main__':
  test()