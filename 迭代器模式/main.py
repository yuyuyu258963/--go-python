
class Iterator:
  def hasNext(self):
    pass
  def next(self):
    pass

class Container:
  def getIterator(self):
    pass

class NameRepository(Container):
  def __init__(self, names):
    self.names = names
  def getIterator(self):
    return NameIterator(self.names)

class NameIterator(Iterator):
  def __init__(self, names):
    self.names = names
    self.index = -1
  def hasNext(self):
    return self.index < len(self.names) - 1
  def next(self):
    if self.hasNext():
      self.index += 1
      return self.names[self.index]
    return None

def IteratorPatternDemo():
  nameRepository = NameRepository(['John', 'Jack', 'Camila', 'Ingrid'])
  iter = nameRepository.getIterator()
  while True:
    if not iter.hasNext():
      break
    name = iter.next()
    print(name)

if __name__ == '__main__':
  IteratorPatternDemo()