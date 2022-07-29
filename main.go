package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	. "go-epub/utils"
	"io/ioutil"
	"strings"
)

//go:generate .\build.bat

type MyMainWindow struct {
	*walk.MainWindow
	edit *walk.TextEdit
	path string
}

func main() {
	var inTE, outTE *walk.TextEdit
	var web *walk.WebView
	mw := &MyMainWindow{}


	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:   "markdownParser",
		MinSize: Size{300, 200},
		Layout:  VBox{},
		// 定义菜单
		MenuItems: []MenuItem{
			Menu{
				Text: "&关于",
				Items: []MenuItem{
					Action{
						Text: "版本",
						OnTriggered: func() {
							walk.MsgBox(mw, "版本", "版本号：0-0-1", walk.MsgBoxIconInformation)
						},
					},
					//...
				},
			},
		},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					//TextEdit{AssignTo: &outTE},
					WebView{
						AssignTo: &web,
						Name:     "wv",
						URL:      ConcatPath( GetCurrentPath(), string("index.html") ), // exe同级的index.html
					},
				},
			},
			PushButton{
				Text: "渲染md",
				OnClicked: func() {
					md := []byte(inTE.Text())
					output := markdown.ToHTML(md, nil, nil)
					outTE.SetText(string(output))
					htmlPath := ConcatPath( GetCurrentPath(), string("index.html") )
					ioutil.WriteFile(htmlPath , output, 0666)
					url := strings.Replace(htmlPath, "/", "\\",0 )
					web.SetURL(url)
				},
			},
			PushButton{
				Text: "播放挤你太美",
				OnClicked: func() {
					playJNTM()
				},
			},
		},
		Functions: map[string]func(args ...interface{}) (interface{}, error){
			"icon": func(args ...interface{}) (interface{}, error) {
				if strings.HasPrefix(args[0].(string), "https") {
					return "check", nil
				}

				return "stop", nil
			},
		},
	}
	MW.Run()
}

func playJNTM(){
	player := MusicEntry{Source: "static/res/jntm.mp3"}
	player.Open()
	player.Play()
}