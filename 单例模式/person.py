import threading
class SingletonType(type):
  _instance_lock = threading.Lock()
  def __call__(cls, *args, **kwargs):
    if not hasattr(cls, "_instance"):
      with cls._instance_lock:
        if not hasattr(cls, "_instance"):
          cls._instance = cls.__call__(*args, **kwargs)
    return cls._instance
class Person(metaclass=SingletonType):
  counter = 1
  def sleep(self):
    self.counter += 1

ywh = Person()
ly = Person()
ly.sleep()
print(ywh.counter)
