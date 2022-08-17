import time

class ChatRoom:
  @classmethod
  def SendMessage(self, user, message):
    print(f"{time.time()} {user.getName()} say: {message} ")

class User:
  def __init__(self, name):
    self.name = name

  def getName(self):
    return self.name
  def setName(self, name):
    self.name = name
  def SendMessage(self, message):
    ChatRoom.SendMessage(self, message)

def test():
  rebort = User("rebort")
  john = User("john")
  rebort.SendMessage("hello john")
  john.SendMessage("hi rebort")

if __name__ == '__main__':
  test()