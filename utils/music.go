package utils

import (
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)


type MusicEntry struct {
	Id         string   //编号
	Name       string   //歌名
	Artist     string   //作者
	Source     string   //位置
	Type       string   //类型
	Filestream *os.File // 文件流
}

func (me *MusicEntry) Open() {
	var err error
	me.Filestream, err = os.Open(me.Source)
	if err != nil {
		log.Fatal(err)
	}
}

func (me *MusicEntry) Play() {
	streamer, format, err := mp3.Decode(me.Filestream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	fmt.Println("music length :", streamer.Len())
	speaker.Play(streamer)
	select {}
}