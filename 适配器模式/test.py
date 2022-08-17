class AdvanceMediaPlayer:
  """
    ## 播放器抽象类
  """
  def play(self, fileName):
    pass

class VlcPlayer(AdvanceMediaPlayer):
  """
    ## Vlc播放器具体实现
  """
  def play(self, fileName):
    print(f"Playing vlc {fileName}")

class Mp4Player(AdvanceMediaPlayer):
  """
    ## MP4播放器具体实现
  """
  def play(self, fileName):
    print(f"Playing mp4 {fileName}")

class MediaPlayer(object):
  def Play(self):
    pass
class MediaAdapter(MediaPlayer):
  def __init__(self):
    self.player = None
  def MediaAdapter(self, audioType):
    if audioType == "vlc":
      self.player = VlcPlayer()
    elif audioType == "mp4":
      self.player = Mp4Player()
  def play(self, fileName):
    self.player.play(fileName)

class AudioPlayer(MediaPlayer):
  def play(self, audioType, fileName):
    if audioType == "mp3":
      print(f"Playing mp3 {fileName}")
    elif audioType == "mp4" or audioType == "vlc":
      player = MediaAdapter()
      player.MediaAdapter(audioType)
      player.play(fileName)

def test():
  player = AudioPlayer()
  player.play("mp3", "beyond the horizon.mp3")
  player.play("mp4", "alone.mp4")
  player.play("vlc", "far far away.vlc")

if __name__ == "__main__":
  test()