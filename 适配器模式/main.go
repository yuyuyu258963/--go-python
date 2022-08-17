package main

import "fmt"

//集成多种方法
type AdvanceMediaPlayer interface {
	PlayMp4(fileName string)
	PlayVlc(fileName string)
}

type VlcPlayer struct {
}

func (p *VlcPlayer) PlayMp4(fileName string) {
}
func (p *VlcPlayer) PlayVlc(fileName string) {
	fmt.Printf("Playing vlc file: %s\n", fileName)
}

type Mp4Player struct {
}

func (p *Mp4Player) PlayVlc(fileName string) {
}
func (p *Mp4Player) PlayMp4(fileName string) {
	fmt.Printf("Playing mp4 file: %s\n", fileName)
}

type MediaPlayer interface {
	Play(audioType string, fileName string)
}

type MediaAdapter struct {
	player AdvanceMediaPlayer
}

func (p *MediaAdapter) initPlayer(audioType string) {
	if audioType == "vlc" {
		p.player = &VlcPlayer{}
	} else if audioType == "mp4" {
		p.player = &Mp4Player{}
	}
}

func (p *MediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		p.player.PlayVlc(fileName)
	} else if audioType == "mp4" {
		p.player.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
}

func (p *AudioPlayer) Play(audioType string, fileName string) {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file: %s\n", fileName)
	} else if audioType == "mp4" || audioType == "vlc" {
		var mediaAdapter MediaAdapter
		mediaAdapter.initPlayer(audioType)
		mediaAdapter.Play(audioType, fileName)
	} else {
		fmt.Printf("Invalid media type: %s\n", audioType)
	}
}

func test() {
	audioPlayer := &AudioPlayer{}
	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone")
	audioPlayer.Play("vlc", "I want it that way")
}

func main() {
	test()
}
