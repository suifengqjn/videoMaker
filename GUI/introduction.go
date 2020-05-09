package GUI

import "github.com/icza/gowut/gwu"

func buildIntroduction(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	label := gwu.NewLabel("核心功能:批量给自动配音 + 自动添加字幕 + 视频合成")
	label.Style().SetFontSize("16")
	label.Style().SetColor(gwu.ClrBlue)
	p.Add(label)

	p.AddVSpace(20)
	link := gwu.NewLink("视频教程", "https://github.com/suifengqjn/videoMaker/blob/master/shipin.md")
	p.Add(link)

	p.AddVSpace(20)
	link = gwu.NewLink("软件地址", "https://github.com/suifengqjn/videoMaker")
	p.Add(link)

	p.AddVSpace(20)
	link = gwu.NewLink("密钥购买", "https://pr.kuaifaka.com/item/3ZUpQ")
	p.Add(link)

	p.AddVSpace(20)
	link = gwu.NewLink("使用文档", "https://www.yuque.com/fengshi-zm9in/bx4gg5")
	p.Add(link)


	p.AddVSpace(20)
	label = gwu.NewLabel("几种模式")
	label.Style().SetFontSize("18")
	p.Add(label)

	arr := []string {
		"模式1: 单视频",
		"模式2: 多视频",
		"模式3: 多图",
		"模式4: 单视频混剪",
		"模式5: 文字转音频",
		"后续增加更多模式...",
	}

	for _, s := range arr {
		p.AddVSpace(5)
		label = gwu.NewLabel(s)
		p.Add(label)
	}

	p.Add(gwu.NewLabel("关注微信公众号，可实时接收软件版本更新信息，后续也会不断在公众号内分享免费软件和自媒体资料"))
	p.Add(gwu.NewLabel("如果你有什么想法，或者想让软件增加什么新功能，都可以在公众号内给我发消息。扫描下方二维码，关注微信公众号"))
	img := gwu.NewImage("公众号", "http://cdn.qiniu.freetop.ren/gzh.jpg")
	img.Style().SetSizePx(200, 200)
	p.Add(img)


	//p.AddVSpace(50)
	//link = gwu.NewLink("常见问题", "https://github.com/suifengqjn/videoWater/blob/master/QREADME.md")
	//p.Add(link)


	return p
}

