package GUI

import "github.com/icza/gowut/gwu"

func buildIntroduction(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	label := gwu.NewLabel("核心功能: 自动配音 + 自动添加字幕")
	label.Style().SetFontSize("16")
	label.Style().SetColor(gwu.ClrBlue)
	p.Add(label)


	p.AddVSpace(20)
	label = gwu.NewLabel("模式距离")
	label.Style().SetFontSize("18")
	p.Add(label)

	arr := []string {
		"模式1: 素材只有一个视频",
		"模式2: 素材一个视频 + 文案",
		"模式2: 素材有两个视频",
	}

	for _, s := range arr {
		p.AddVSpace(5)
		label = gwu.NewLabel(s)
		p.Add(label)
	}

	return p
}

