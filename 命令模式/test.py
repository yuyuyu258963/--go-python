class Command(object):
  """
    ## 命令的接口
  """
  def Execute(self):
    pass

class MoveCommand(Command):
  """
    ## 移动命令
  """
  def __init__(self, x, y):
    self.x = x
    self.y = y
  def Execute(self):
    print(f"Move to {self.x},{self.y}")


class AttackCommand(Command):
  """
    ## 攻击命名
  """
  def __init__(self, skill):
    self.skill = skill
  def Execute(self):
    print(f"Attack with {self.skill}")

def AddCommand(action):
  """
    ## 添加命令
  """
  if action == "attack":
    return AttackCommand("crush")
  else :
    return MoveCommand(1,2)

def test():
  commandList = []
  commandList.append(AddCommand("attack"))
  commandList.append(AddCommand("move"))
  commandList.append(AddCommand("attack"))
  for command in commandList:
    command.Execute()

if __name__ == '__main__':
  test()
