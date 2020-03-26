// Copyright (C) 2013 Andras Belicza. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// A Gowut "Showcase of Features" application.

package GUI

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"log"
)


var tempEvent gwu.Event
// plural returns an empty string if i is equal to 1,
// "s" otherwise.
func plural(i int) string {
	if i == 1 {
		return ""
	}
	return "s"
}


func buildCollectDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	p.AddVSpace(20)
	link := gwu.NewLink("视频采集器", "https://github.com/suifengqjn/videoCollector")
	p.Add(link)
	return p
}

func buildEditorDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()
	p.AddVSpace(20)
	link := gwu.NewLink("视频剪辑器", "https://github.com/suifengqjn/videoWater")
	p.Add(link)
	return p
}

type demo struct {
	link      gwu.Label
	buildFunc func(gwu.Event) gwu.Comp
	comp      gwu.Comp // Lazily initialized demo comp
}
type pdemo *demo

var extraHeadHTMLs []string

func buildHome(sess gwu.Session) {
	win := gwu.NewWindow("show", "生成器")
	for _, headHTML := range extraHeadHTMLs {
		win.AddHeadHTML(headHTML)
	}

	win.Style().SetFullSize()
	win.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETypeWinLoad:
			//log.Println("LOADING window:", e.Src().ID())
		case gwu.ETypeWinUnload:
			//log.Println("UNLOADING window:", e.Src().ID())
		}
	}, gwu.ETypeWinLoad, gwu.ETypeWinUnload)

	hiddenPan := gwu.NewNaturalPanel()
	sess.SetAttr("hiddenPan", hiddenPan)

	header := gwu.NewHorizontalPanel()
	header.Style().SetFullWidth().SetBorderBottom2(2, gwu.BrdStyleSolid, "#cccccc")
	title := gwu.NewLink("AI全自动原创视频生成器", win.Name())
	title.SetTarget("")
	title.Style().SetColor(gwu.ClrBlue).SetFontWeight(gwu.FontWeightBold).SetFontSize("120%").Set("text-decoration", "none")
	header.Add(title)


	//header.AddHSpace(200)
	//startBtn := gwu.NewButton("开始处理")
	//startBtn.Style().SetFontSize("18")
	//startBtn.Style().SetColor(gwu.ClrBlue)
	//
	//startBtn.AddEHandlerFunc(func(e gwu.Event) {
	//	if account.AppAccount.IsActive() == false {
	//		titleLabel.SetText("请先在激活页面激活软件")
	//		titleLabel.Style().SetColor("red")
	//		e.MarkDirty(titleLabel)
	//	}
	//	go app.Engine.DoFactory()
	//
	//}, gwu.ETypeClick)
	//header.Add(startBtn)
	//
	//header.AddHSpace(50)
	//clearBtn := gwu.NewButton("清空配置")
	//clearBtn.Style().SetFontSize("18")
	//clearBtn.Style().SetColor(gwu.ClrBlue)
	//clearBtn.AddEHandlerFunc(func(e gwu.Event) {
	//
	//}, gwu.ETypeClick)
	//header.Add(clearBtn)

	header.AddHConsumer()

	header.AddHSpace(100)
	setNoWrap(header)
	win.Add(header)

	content := gwu.NewHorizontalPanel()
	content.SetCellPadding(1)
	content.SetVAlign(gwu.VATop)
	content.Style().SetFullSize()

	demoWrapper := gwu.NewPanel()
	demoWrapper.Style().SetPaddingLeftPx(5)
	demoWrapper.AddVSpace(10)
	demoTitle := gwu.NewLabel("")
	demoTitle.Style().SetFontWeight(gwu.FontWeightBold).SetFontSize("100%")
	demoWrapper.Add(demoTitle)
	demoWrapper.AddVSpace(10)

	links := gwu.NewPanel()
	links.SetCellPadding(1)
	links.Style().SetPaddingRightPx(5)

	demos := make(map[string]pdemo)
	var selDemo pdemo

	selectDemo := func(d pdemo, e gwu.Event) {
		if selDemo != nil {
			selDemo.link.Style().SetBackground("")
			if e != nil {
				e.MarkDirty(selDemo.link)
			}
			demoWrapper.Remove(selDemo.comp)
		}
		selDemo = d
		d.link.Style().SetBackground("#88ff88")
		demoTitle.SetText(d.link.Text())
		if d.comp == nil {
			d.comp = d.buildFunc(e)
		}
		demoWrapper.Add(d.comp)
		if e != nil {
			e.MarkDirty(d.link, demoWrapper)
		}
	}

	createDemo := func(name string, buildFunc func(gwu.Event) gwu.Comp) pdemo {
		link := gwu.NewLabel(name)
		link.Style().SetFullWidth().SetCursor(gwu.CursorPointer).SetDisplay(gwu.DisplayBlock).SetColor(gwu.ClrBlue)
		demo := &demo{link: link, buildFunc: buildFunc}
		link.AddEHandlerFunc(func(e gwu.Event) {
			selectDemo(demo, e)
			if tempEvent == nil {
				tempEvent = e
			}
		}, gwu.ETypeClick)
		links.Add(link)
		demos[name] = demo
		return demo
	}

	links.Style().SetFullHeight().SetBorderRight2(2, gwu.BrdStyleSolid, "#cccccc")
	links.AddVSpace(5)
	homeDemo := createDemo("介绍", buildIntroduction)
	selectDemo(homeDemo, nil)
	links.AddVSpace(5)


	//=============================================//
	l0 := gwu.NewLabel("账户")
	l0.Style().SetFontWeight(gwu.FontWeightBold)
	links.Add(l0)
	createDemo("激活", buildActiveUI)
	createDemo("参数", buildPlatform)
	//=============================================//


	links.AddVSpace(10)
	l := gwu.NewLabel("生成原创")
	l.Style().SetFontWeight(gwu.FontWeightBold)
	links.Add(l)
	wei := createDemo("配置", buildMainShowUI)
	selectDemo(wei, nil)
	//=============================================//


	links.AddVSpace(20)
	l = gwu.NewLabel("其他软件")
	l.Style().SetFontWeight(gwu.FontWeightBold)
	links.Add(l)
	createDemo("视频采集器", buildCollectDemo)
	createDemo("视频剪辑器", buildEditorDemo)


	links.AddVConsumer()
	setNoWrap(links)
	content.Add(links)
	content.Add(demoWrapper)
	content.CellFmt(demoWrapper).Style().SetFullWidth()

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

	footer := gwu.NewHorizontalPanel()
	footer.Style().SetFullWidth().SetBorderTop2(2, gwu.BrdStyleSolid, "#cccccc")
	footer.Add(hiddenPan)
	footer.AddHConsumer()
	l = gwu.NewLabel("自媒体软件工作室")
	l.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(l)
	footer.AddHSpace(10)
	link := gwu.NewLink("点击进入", "http://freetop.ren")
	link.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(link)
	setNoWrap(footer)
	win.Add(footer)

	sess.AddWin(win)
}

// setNoWrap sets WhiteSpaceNowrap to all children of the specified panel.
func setNoWrap(panel gwu.Panel) {
	count := panel.CompsCount()
	for i := count - 1; i >= 0; i-- {
		panel.CompAt(i).Style().SetWhiteSpace(gwu.WhiteSpaceNowrap)
	}
}

// SessHandler is our session handler to build the showcases window.
type sessHandler struct{}

func (h sessHandler) Created(s gwu.Session) {
	buildHome(s)
}

func (h sessHandler) Removed(s gwu.Session) {}

// StartServer creates and starts the Gowut GUI server.
func StartServer(appName, addr string, autoOpen bool) {
	// Create GUI server
	server := gwu.NewServer(appName, addr)
	for _, headHTML := range extraHeadHTMLs {
		server.AddRootHeadHTML(headHTML)
	}
	server.SetText("视频伪原创剪辑器")

	server.AddSessCreatorName("show", "点击进入")
	server.AddSHandler(sessHandler{})
	// Just for the demo: Add an extra "Gowut-Server" header to all responses holding the Gowut version
	server.SetHeaders(map[string][]string{
		"Gowut-Server": {gwu.GowutVersion},
	})

	// Start GUI server
	var openWins []string
	if autoOpen {
		openWins = []string{"show"}
	}
	fmt.Println("软件已经启动，请不要退出，可以最小化本窗口")
	if err := server.Start(openWins...); err != nil {
		log.Println("Error: Cound not start GUI server:", err)
		return
	}
}
