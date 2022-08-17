class Packing(object):
  def Pack(self):
    pass

class Wrapper(Packing):
  def Pack(self):
    return "Pack :: Wrapper"

class Bottle(Packing):
  def Pack(self):
    return "Pack :: Bottle"