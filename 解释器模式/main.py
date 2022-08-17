class Context(object):
  def __init__(self, action ,context):
    self.action = action
    self.context = context

class InterPreter(object):
  """
    ## 翻译类的接口
  """
  def InterPreter(self, context):
    pass

class MusicInterpreter(InterPreter):
  """
    ## 音乐翻译类
  """
  def InterPreter(self, context):
    print(f"{context.action} 中的 {context.context} 很好听")    

class MovieInterpreter(InterPreter):
  """
    ## 电影翻译类
  """
  def InterPreter(self, context):
    print(f"{context.action} 中的 {context.context} 很好看")

def test():
  actionList = [
    Context("music", "高音"),
    Context("music", "低音"),
    Context("movie", "电影"),
  ]
  for c in actionList:
    if c.action == "music":
      MusicInterpreter().InterPreter(c)
    elif c.action == "movie":
      MovieInterpreter().InterPreter(c)

if __name__ == '__main__':
  test()