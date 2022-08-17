class ReadFile:
  """
    ## 各种文件的抽象类型
  """
  def Read(self):
    pass
  def Accept(self):
    pass

class ReadPdfFile(ReadFile):
  def Read(self):
    print("Read PDF File")
  def Accept(self, visitor):
    visitor.VisitPdfFile(self)

class ReadTxtFile(ReadFile):
  def Read(self):
    print("Read TXT File")
  def Accept(self, visitor):
    visitor.VisitTxtFile(self)

class Visitor:
  """
    ## 访问者抽象类
  """
  def VistorPdfFile(self, r):
    pass
  def VistorTxtFile(self, r):
    pass

class ExactFile(Visitor):
  """
    ## 读取文件的类
  """
  def VisitPdfFile(self, file):
    print("读取pdf文件")
  def VisitTxtFile(self, file):
    print("读取txt文件")

class CompressionFile(Visitor):
  """
    ## 压缩文件的类
  """
  def VisitPdfFile(self, file):
    print("压缩pdf文件")
  def VisitTxtFile(self, file):
    print("压缩txt文件")

def test():
  fileList = [
    ReadPdfFile(),
    ReadTxtFile(),
    ReadPdfFile(),
  ]
  extract = ExactFile()
  compress = CompressionFile()
  for file in fileList:
    file.Accept(extract)
  for file in fileList:
    file.Accept(compress)

if __name__ == '__main__':
  test()