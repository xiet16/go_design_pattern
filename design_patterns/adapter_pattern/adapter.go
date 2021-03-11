package adapter_pattern

import "fmt"

/*

适配器模式：将一个类适配成用户所需要的接口
理解：1、该类的原本的功能还在，只是扩展    2、扩展类的功能应该属于同一类功能   3、使用继承或依赖的方式，不然就是组合模式了

*/

/*
 例子说明：
  MediaPlay 是一个接口， AudioPlay 继承了MediaPlay , 这个类只能播放MP3音乐
  AdvancedPlay 是一个接口， VlcPlay 和 Mp4Play 继承了AdvancedPlay ， 它们可以播放MP4 和vlc
  现在想让AudioPlay 可以播放Mp4 和vlc
  做法是增加一个AdapterPlay , 去继承AvancedPlay ,然后在AudioPlay 中使用
*/

type MediaPlay interface {
	Play(fileType string, file string)
}
type AdvancedPlay interface {
	PlayVlc(fileType string, file string)
	PlayMp4(fileType string, file string)
}

type AudioPlay struct {
	AdvancedPlay AdvancedPlay
}

func (p *AudioPlay) Play(fileType string, file string) {
	if fileType == "mp3" {
		fmt.Println("AudioPlay play mp3")
	} else if fileType == "mp4" {
		p.AdvancedPlay.PlayMp4(fileType, file)
	} else if fileType == "vlc" {
		p.AdvancedPlay.PlayVlc(fileType, file)
	}
}

type VlcPlay struct{}
type Mp4Play struct{}
type AdapterPlay struct{}

func (p *VlcPlay) PlayVlc(fileType string, file string) {
	fmt.Println("vlcplay play vlc")
}

func (p *VlcPlay) PlayMp4(fileType string, file string) {
	fmt.Println("vlcplay play mp4")
}

func (p *Mp4Play) PlayVlc(fileType string, file string) {
	fmt.Println("Mp4Play play vlc")
}

func (p *Mp4Play) PlayMp4(fileType string, file string) {
	fmt.Println("Mp4Play play mp4")
}

func (p *AdapterPlay) PlayVlc(fileType string, file string) {
	fmt.Println("AdapterPlay play vlc")
}

func (p *AdapterPlay) PlayMp4(fileType string, file string) {
	fmt.Println("AdapterPlay play mp4")
}
