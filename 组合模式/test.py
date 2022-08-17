class ComponentBases:
  """
    ## 文件抽象类
  """
  def __init__(self, fileName):
    self.fileName = fileName
    pass
  def setFileName(self, fileName):
    self.fileName = fileName
  def display(self, Separator:str):
    pass

class FileNode(ComponentBases):
  """
    ## 文件类
  """
  def __init__(self, fileName):
    super().__init__(fileName)
  def display(self, Separator:str):
    print(f"{Separator}FileNode: {self.fileName}")

class directoryFileNode(ComponentBases):
  """
    ## 目录类
  """
  def __init__(self, fileName):
    super().__init__(fileName)
    self.children = []
  def add(self, obj):
    self.children.append(obj)
  def display(self, Separator:str):
    print(f"{Separator}{self.fileName}文件内容为:")
    for child in self.children:
      child.display(Separator+Separator)

def test():
  root = directoryFileNode("root")
  root.add(FileNode("file1"))
  root.add(FileNode("file2"))
  file2 = directoryFileNode("root2")
  file2.add(FileNode("file3"))
  root.add(file2)
  root.display(">>")

if __name__ == '__main__':
  test()