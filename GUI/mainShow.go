package GUI

import (
	"github.com/icza/gowut/gwu"
	"myProject/videoMaker/account"
	"myProject/videoMaker/app"
	"myProject/videoMaker/common"
	"strings"
)

var GRAY = "#808080"
var titleLabel gwu.Label
var defaultShowString = "原创参数配置, 需要剪辑的视频放在video目录下,首次使用一定要看教程"

func buildMainShowUI(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	buildUI(p)

	return p
}

func buildUI(p gwu.Panel) {

	titleLabel = gwu.NewLabel(defaultShowString)
	titleLabel.Style().SetColor(gwu.ClrBlue)

	p.Add(titleLabel)

	//p.AddVSpace(10)
	//buildFourCutFront(p)
	//
	//p.AddVSpace(10)
	//buildFiveCutBack(p)
	//
	//p.AddVSpace(10)
	//buildSevenClearWater(p)
	//
	//p.AddVSpace(10)
	//buildSevenClearWater2(p)

	//-------//
	p.AddVSpace(10)
	buildSrtUI(p)

	p.AddVSpace(10)
	buildComposite(p)

	p.AddVSpace(10)
	buildSubTitle(p)
	//-------//

	p.AddVSpace(10)
	buildTwelveTextWater(p)

	p.AddVSpace(10)
	buildTwelveTextWater2(p)

	p.AddVSpace(10)
	buildThirteenImageWater(p)

	p.AddVSpace(10)
	buildFiveteenBgm(p)

	p.AddVSpace(10)
	buildSixteenHeadEnd(p)

	p.AddVSpace(25)
	buildBottomBtn(p)
}

func buildFourCutFront(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	cutHeadCb = gwu.NewCheckBox("剪去片头")
	cutHeadCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.CutFront.Switch = cutHeadCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(cutHeadCb)

	row.AddHSpace(18)

	cutHeadTb = gwu.NewTextBox("")
	cutHeadTb.SetMaxLength(10)
	cutHeadTb.Style().SetWidthPx(50)
	cutHeadTb.AddSyncOnETypes(gwu.ETypeKeyUp)

	cutHeadTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.CutFront.Value = common.IntValue(cutHeadTb.Text())

	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(cutHeadTb)

	row.AddHSpace(18)
	descLabel := gwu.NewLabel("如: 7 ,剪去片头7秒")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)

}

func buildFiveCutBack(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	cutBackCb = gwu.NewCheckBox("剪去片尾")
	cutBackCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.CutBack.Switch = cutBackCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(cutBackCb)
	row.AddHSpace(18)

	cutBackTb = gwu.NewTextBox("")
	cutBackTb.SetMaxLength(10)
	cutBackTb.Style().SetWidthPx(50)
	cutBackTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	cutBackTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.CutBack.Value = common.IntValue(cutBackTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(cutBackTb)

	row.AddHSpace(18)
	descLabel := gwu.NewLabel("如: 7 ,剪去片尾7秒")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)

}

func buildSevenClearWater(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	ClearWaterCb = gwu.NewCheckBox("去除水印")
	ClearWaterCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater.Switch = ClearWaterCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(ClearWaterCb)
	row.AddHSpace(18)

	xl := gwu.NewLabel("x:")
	row.Add(xl)
	ClearWaterTbx = gwu.NewTextBox("")
	ClearWaterTbx.SetMaxLength(5)
	ClearWaterTbx.Style().SetWidthPx(50)
	ClearWaterTbx.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWaterTbx.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.ClearWater.X = common.IntValue(ClearWaterTbx.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWaterTbx)

	row.AddHSpace(5)
	yl := gwu.NewLabel("y:")
	row.Add(yl)
	ClearWaterTby = gwu.NewTextBox("")
	ClearWaterTby.Style().SetBorder("1")
	ClearWaterTby.SetMaxLength(5)
	ClearWaterTby.Style().SetWidthPx(50)
	ClearWaterTby.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWaterTby.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater.Y = common.IntValue(ClearWaterTby.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWaterTby)

	row.AddHSpace(5)
	wl := gwu.NewLabel("w:")
	row.Add(wl)
	ClearWaterTbw = gwu.NewTextBox("")
	ClearWaterTbw.Style().SetBorder("1")
	ClearWaterTbw.SetMaxLength(5)
	ClearWaterTbw.Style().SetWidthPx(50)
	ClearWaterTbw.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWaterTbw.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater.W = common.IntValue(ClearWaterTbw.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWaterTbw)

	row.AddHSpace(5)
	hl := gwu.NewLabel("h:")
	row.Add(hl)
	ClearWaterTbh = gwu.NewTextBox("")
	ClearWaterTbh.Style().SetBorder("1")
	ClearWaterTbh.SetMaxLength(5)
	ClearWaterTbh.Style().SetWidthPx(50)
	ClearWaterTbh.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWaterTbh.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.ClearWater.H = common.IntValue(ClearWaterTbh.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWaterTbh)

	row.AddHSpace(18)
	descLabel := gwu.NewLabel("x,y 起始坐标，w 宽度，h 高度")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)

}

func buildSevenClearWater2(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	ClearWater2Cb = gwu.NewCheckBox("去除水印2")
	ClearWater2Cb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.ClearWater1.Switch = ClearWater2Cb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(ClearWater2Cb)
	row.AddHSpace(10)

	xl := gwu.NewLabel("x:")
	row.Add(xl)
	ClearWater2Tbx = gwu.NewTextBox("")
	ClearWater2Tbx.SetMaxLength(5)
	ClearWater2Tbx.Style().SetWidthPx(50)
	ClearWater2Tbx.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWater2Tbx.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater1.X = common.IntValue(ClearWater2Tbx.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWater2Tbx)

	row.AddHSpace(5)
	yl := gwu.NewLabel("y:")
	row.Add(yl)
	ClearWater2Tby = gwu.NewTextBox("")
	ClearWater2Tby.Style().SetBorder("1")
	ClearWater2Tby.SetMaxLength(5)
	ClearWater2Tby.Style().SetWidthPx(50)
	ClearWater2Tby.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWater2Tby.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater1.Y = common.IntValue(ClearWater2Tby.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWater2Tby)

	row.AddHSpace(5)
	wl := gwu.NewLabel("w:")
	row.Add(wl)
	ClearWater2Tbw = gwu.NewTextBox("")
	ClearWater2Tbw.Style().SetBorder("1")
	ClearWater2Tbw.SetMaxLength(5)
	ClearWater2Tbw.Style().SetWidthPx(50)
	ClearWater2Tbw.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWater2Tbw.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.ClearWater1.W = common.IntValue(ClearWater2Tbw.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWater2Tbw)

	row.AddHSpace(5)
	hl := gwu.NewLabel("h:")
	row.Add(hl)
	ClearWater2Tbh = gwu.NewTextBox("")
	ClearWater2Tbh.Style().SetBorder("1")
	ClearWater2Tbh.SetMaxLength(5)
	ClearWater2Tbh.Style().SetWidthPx(50)
	ClearWater2Tbh.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWater2Tbh.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.ClearWater1.H = common.IntValue(ClearWater2Tbh.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWater2Tbh)

	row.AddHSpace(18)
	descLabel := gwu.NewLabel("x,y 起始坐标，w 宽度，h 高度")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)

}

//-----//
func buildSrtUI(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	SrtCb = gwu.NewCheckBox("提取字幕")
	SrtCb.Style().SetColor("#FF6633")
	SrtCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.ExtractSubtitles.Switch = SrtCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(SrtCb)
	p.Add(row)
}

func buildComposite(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	CompositeCb = gwu.NewCheckBox("视频合成")
	CompositeCb.Style().SetColor("#FF6633")
	CompositeCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Composite.Switch = CompositeCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row.Add(CompositeCb)


	row.AddHSpace(10)
	row.Add(gwu.NewLabel("模式:"))
	CompositeLb = gwu.NewListBox(common.CompleteStyle)
	CompositeLb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Composite.Style = CompositeLb.SelectedIdx()
	}, gwu.ETypeChange)
	row.Add(CompositeLb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("配音员:"))
	SpeecherLb = gwu.NewListBox(common.Voiceover)
	common.AppConfig.Composite.Voice = "Xiaoyun"
	SpeecherLb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Composite.Voice = common.VoiceoverMap[SpeecherLb.SelectedValue()]

	}, gwu.ETypeChange)
	row.Add(SpeecherLb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("音量:"))
	VolumeTB = gwu.NewTextBox("50")
	VolumeTB.SetMaxLength(3)
	VolumeTB.Style().SetWidthPx(50)
	VolumeTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	common.AppConfig.Composite.Volume = 50
	VolumeTB.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.Composite.Volume =  common.IntValue(VolumeTB.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(VolumeTB)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("语速:"))
	SpeexhRate = gwu.NewTextBox("0")
	SpeexhRate.SetMaxLength(3)
	SpeexhRate.Style().SetWidthPx(50)
	SpeexhRate.AddSyncOnETypes(gwu.ETypeKeyUp)
	SpeexhRate.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Composite.SpeechRate = common.IntValue(SpeexhRate.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SpeexhRate)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("语调:"))
	PitchRate = gwu.NewTextBox("0")
	PitchRate.SetMaxLength(3)
	PitchRate.Style().SetWidthPx(50)
	PitchRate.AddSyncOnETypes(gwu.ETypeKeyUp)
	PitchRate.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Composite.PitchRate = common.IntValue(PitchRate.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PitchRate)

	row.AddHSpace(20)
	link0 := gwu.NewLink("播音人","https://github.com/suifengqjn/videoMaker/blob/master/voicer.md")
	row.Add(link0)

	p.Add(row)

	p.AddVSpace(10)
	row2 := gwu.NewHorizontalPanel()
	row2.AddHSpace(100)
	desc := gwu.NewLabel("音量,范围是0~100，默认50. 语速和语调范围是-500~500，默认是0")
	desc.Style().SetColor(GRAY)
	row2.Add(desc)

	row2.AddHSpace(30)
	link := gwu.NewLink("颜色值选取","http://cha.buyiju.com/tool/color.html")
	row2.Add(link)

	p.Add(row2)
}

func buildSubTitle(p gwu.Panel)  {

	row := gwu.NewHorizontalPanel()
	SubTitleCb = gwu.NewCheckBox("字幕属性")
	SubTitleCb.Style().SetColor("#FF6633")
	SubTitleCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.Switch = SubTitleCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row.Add(SubTitleCb)


	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字体大小:"))
	SubTitleFontSizeTb = gwu.NewTextBox("")
	SubTitleFontSizeTb.SetMaxLength(3)
	SubTitleFontSizeTb.Style().SetWidthPx(50)
	SubTitleFontSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontSizeTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.Subtitles.FontSize =  common.IntValue(SubTitleFontSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SubTitleFontSizeTb)


	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字体颜色:"))
	SubTitleFontColorTb = gwu.NewTextBox("")
	SubTitleFontColorTb.SetMaxLength(7)
	SubTitleFontColorTb.Style().SetWidthPx(50)
	SubTitleFontColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.FontColor = SubTitleFontColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SubTitleFontColorTb)


	row.AddHSpace(10)
	row.Add(gwu.NewLabel("底部距离:"))
	SubTitleMarginVTb = gwu.NewTextBox("")
	SubTitleMarginVTb.SetMaxLength(3)
	SubTitleMarginVTb.Style().SetWidthPx(50)
	SubTitleMarginVTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleMarginVTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.MarginV = common.IntValue(SubTitleMarginVTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SubTitleMarginVTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("背景条颜色:"))
	BjColorTb = gwu.NewTextBox("")
	BjColorTb.SetMaxLength(7)
	BjColorTb.Style().SetWidthPx(50)
	BjColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	BjColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.BjColor = BjColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(BjColorTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("背景条透明度:"))
	BjAlphaTb = gwu.NewTextBox("")
	BjAlphaTb.SetMaxLength(3)
	BjAlphaTb.Style().SetWidthPx(50)
	BjAlphaTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	BjAlphaTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.BjAlpha = common.IntValue(BjAlphaTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(BjAlphaTb)

	row.AddHSpace(10)
	desc := gwu.NewLabel("透明度 0-10, 0表示完全透明, 10不透明")
	desc.Style().SetColor(GRAY)
	row.Add(desc)
	p.Add(row)


	p.AddVSpace(10)
	row2 := gwu.NewHorizontalPanel()
	CoverBjCb = gwu.NewCheckBox("遮盖原字幕")
	CoverBjCb.Style().SetColor("#FF6633")
	CoverBjCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.CoverBj = CoverBjCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row2.Add(CoverBjCb)

	row2.AddHSpace(10)
	row2.Add(gwu.NewLabel("遮盖高度:"))
	CoverHTb = gwu.NewTextBox("")
	CoverHTb.SetMaxLength(3)
	CoverHTb.Style().SetWidthPx(50)
	CoverHTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	CoverHTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.CoverH = common.IntValue(CoverHTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(CoverHTb)

	row2.AddHSpace(10)
	desc = gwu.NewLabel("高度不填,会自动识别高度")
	desc.Style().SetColor(GRAY)
	row2.Add(desc)

	p.Add(row2)

}

//文字水印
func buildTwelveTextWater(p gwu.Panel) {
	row1 := gwu.NewHorizontalPanel()
	WaterTextCb = gwu.NewCheckBox("文字水印")
	WaterTextCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.WaterText.Switch = WaterTextCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row1.Add(WaterTextCb)

	row1.AddHSpace(18)
	content := gwu.NewLabel("内容:")
	row1.Add(content)
	WaterTextTb = gwu.NewTextBox("")
	WaterTextTb.SetToolTip("123123")
	WaterTextTb.SetMaxLength(500)
	WaterTextTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterTextTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.WaterText.Content = WaterTextTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(WaterTextTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("大小:")
	row1.Add(content)
	WaterTextSizeTb = gwu.NewTextBox("")
	WaterTextSizeTb.SetMaxLength(10)
	WaterTextSizeTb.Style().SetWidthPx(50)
	WaterTextSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterTextSizeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.WaterText.Size = common.IntValue(WaterTextSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(WaterTextSizeTb)
	WaterTextSizeTb.Style().SetWidthPx(50)

	row1.AddHSpace(10)
	content = gwu.NewLabel("颜色:")

	row1.Add(content)
	WaterTextColorLb = gwu.NewListBox(common.TextColors)
	WaterTextColorLb.AddEHandlerFunc(func(e gwu.Event) {
		//fmt.Println(WaterTextColorLb.SelectedIdx(),WaterTextColorLb.SelectedValue())
		if WaterTextColorLb.SelectedIdx() > 0 {
			common.AppConfig.WaterText.Color = WaterTextColorLb.SelectedValue()
		}

	}, gwu.ETypeChange)
	row1.Add(WaterTextColorLb)

	//font
	row1.AddHSpace(10)
	label := gwu.NewLabel("选择字体:")
	row1.Add(label)
	fonts, keys := common.LoadFonts()
	arr := []string{"默认"}
	arr = append(arr, keys...)
	WaterTextFontLb = gwu.NewListBox(arr)
	WaterTextFontLb.AddEHandlerFunc(func(e gwu.Event) {
		if WaterTextFontLb.SelectedIdx() > 0 {
			common.AppConfig.WaterText.Path = fonts[WaterTextFontLb.SelectedValue()]
		} else {
			common.AppConfig.WaterText.Path = ""
		}

	}, gwu.ETypeChange)
	row1.Add(WaterTextFontLb)

	p.Add(row1)

	p.AddVSpace(5)
	row2 := gwu.NewHorizontalPanel()
	row2.AddHSpace(100)
	content = gwu.NewLabel("样式:")
	row2.Add(content)
	WaterTextStyleLb = gwu.NewListBox(common.WaterStyle)
	WaterTextStyleLb.AddEHandlerFunc(func(e gwu.Event) {
		//fmt.Println(WaterTextStyleLb.SelectedIdx(),WaterTextStyleLb.SelectedValue())
		for k, v := range common.WaterStyleMap {
			if v == strings.TrimSpace(WaterTextStyleLb.SelectedValue()) {
				common.AppConfig.WaterText.Style = k
			}
		}

	}, gwu.ETypeChange)
	row2.Add(WaterTextStyleLb)

	row2.AddHSpace(10)
	content = gwu.NewLabel("水平间距:")
	row2.Add(content)
	WaterTextHSpanTb = gwu.NewTextBox("")
	WaterTextHSpanTb.SetMaxLength(10)
	WaterTextHSpanTb.Style().SetWidthPx(50)
	WaterTextHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterTextHSpanTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.WaterText.Sp1 = common.IntValue(WaterTextHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(WaterTextHSpanTb)

	row2.AddHSpace(10)
	content = gwu.NewLabel("垂直间距:")
	row2.Add(content)
	WaterTextVSpanTb = gwu.NewTextBox("")
	WaterTextVSpanTb.SetMaxLength(10)
	WaterTextVSpanTb.Style().SetWidthPx(50)
	WaterTextVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterTextVSpanTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.WaterText.Sp2 = common.IntValue(WaterTextVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(WaterTextVSpanTb)
	p.Add(row2)
}

//滚动文字
func buildTwelveTextWater2(p gwu.Panel) {
	row1 := gwu.NewHorizontalPanel()
	RunWaterTextCb = gwu.NewCheckBox("滚动文字")
	RunWaterTextCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.RunWaterText.Switch = RunWaterTextCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row1.Add(RunWaterTextCb)

	row1.AddHSpace(18)
	content := gwu.NewLabel("内容:")
	row1.Add(content)
	RunWaterTextTb = gwu.NewTextBox("")
	RunWaterTextTb.SetMaxLength(500)
	RunWaterTextTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	RunWaterTextTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.RunWaterText.Content = RunWaterTextTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(RunWaterTextTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("大小:")
	row1.Add(content)
	RunWaterTextSizeTb = gwu.NewTextBox("")
	RunWaterTextSizeTb.SetMaxLength(10)
	RunWaterTextSizeTb.Style().SetWidthPx(50)
	RunWaterTextSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	RunWaterTextSizeTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.RunWaterText.Size = common.IntValue(RunWaterTextSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(RunWaterTextSizeTb)

	row1.AddHSpace(10)
	content = gwu.NewLabel("颜色:")
	row1.Add(content)
	RunWaterTextColorLb = gwu.NewListBox(common.TextColors)
	RunWaterTextColorLb.AddEHandlerFunc(func(e gwu.Event) {
		if RunWaterTextColorLb.SelectedIdx() > 0 {
			common.AppConfig.RunWaterText.Color = RunWaterTextColorLb.SelectedValue()
		}
	}, gwu.ETypeChange)
	row1.Add(RunWaterTextColorLb)

	row1.AddHSpace(10)
	label := gwu.NewLabel("选择字体:")
	row1.Add(label)

	fonts, keys := common.LoadFonts()
	arr := []string{"默认"}
	arr = append(arr, keys...)
	RunWaterTextFontLb = gwu.NewListBox(arr)
	RunWaterTextFontLb.AddEHandlerFunc(func(e gwu.Event) {
		if RunWaterTextFontLb.SelectedIdx() > 0 {
			common.AppConfig.RunWaterText.Path = fonts[RunWaterTextFontLb.SelectedValue()]
		} else {
			common.AppConfig.RunWaterText.Path = ""
		}

	}, gwu.ETypeChange)
	row1.Add(RunWaterTextFontLb)

	p.Add(row1)

	p.AddVSpace(5)
	row2 := gwu.NewHorizontalPanel()
	row2.AddHSpace(100)
	content = gwu.NewLabel("样式:")
	row2.Add(content)
	RunWaterTextStyleLb = gwu.NewListBox([]string{"未选择", "底部", "顶部"})
	RunWaterTextStyleLb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.RunWaterText.IsTop = RunWaterTextStyleLb.SelectedIdx()

	}, gwu.ETypeChange)
	row2.Add(RunWaterTextStyleLb)

	row2.AddHSpace(10)
	content = gwu.NewLabel("方向:")
	row2.Add(content)
	RunWaterTextDirectionLb = gwu.NewListBox([]string{"未选择", "从右到左", "从左到右"})
	RunWaterTextDirectionLb.AddEHandlerFunc(func(e gwu.Event) {


	}, gwu.ETypeChange)
	row2.Add(RunWaterTextDirectionLb)

	row2.AddHSpace(10)
	content = gwu.NewLabel("间距:")
	row2.Add(content)
	RunWaterTextVSpanTb = gwu.NewTextBox("")
	RunWaterTextVSpanTb.SetMaxLength(10)
	RunWaterTextVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	RunWaterTextVSpanTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.RunWaterText.Sp = common.IntValue(RunWaterTextVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(RunWaterTextVSpanTb)
	p.Add(row2)
}

//图片水印
func buildThirteenImageWater(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	WaterImageCb = gwu.NewCheckBox("图片水印")
	WaterImageCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.WaterImage.Switch = WaterImageCb.State()
	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row.Add(WaterImageCb)

	row.AddHSpace(18)
	image := gwu.NewLabel("图片名称:")
	row.Add(image)

	images, keys := common.LoadImages()

	arr := []string{"未选择"}
	arr = append(arr, keys...)
	WaterImageLb = gwu.NewListBox(arr)
	WaterImageLb.AddEHandlerFunc(func(e gwu.Event) {
		if WaterImageLb.SelectedIdx() > 0 {
			common.AppConfig.WaterImage.Path = images[WaterImageLb.SelectedValue()]
		} else {
			common.AppConfig.WaterImage.Path = ""
		}
		//fmt.Println(common.AppConfig.WaterImage.Path)
	}, gwu.ETypeChange)
	row.Add(WaterImageLb)

	content := gwu.NewLabel("样式:")
	row.Add(content)
	WaterImageStyleLb = gwu.NewListBox(common.WaterStyle)
	WaterImageStyleLb.AddEHandlerFunc(func(e gwu.Event) {
		for k, v := range common.WaterStyleMap {
			if v == strings.TrimSpace(WaterImageStyleLb.SelectedValue()) {
				common.AppConfig.WaterImage.Style = k
			}
		}
	}, gwu.ETypeChange)
	row.Add(WaterImageStyleLb)

	row.AddHSpace(18)
	content = gwu.NewLabel("水平间距:")
	row.Add(content)
	WaterImageHSpanTb = gwu.NewTextBox("")
	WaterImageHSpanTb.SetMaxLength(10)
	WaterImageHSpanTb.Style().SetWidthPx(50)
	WaterImageHSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterImageHSpanTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.WaterImage.Sp1 = common.IntValue(WaterImageHSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(WaterImageHSpanTb)

	row.AddHSpace(18)
	content = gwu.NewLabel("垂直间距:")
	row.Add(content)
	WaterImageVSpanTb = gwu.NewTextBox("")
	WaterImageVSpanTb.SetMaxLength(10)
	WaterImageVSpanTb.Style().SetWidthPx(50)
	WaterImageVSpanTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterImageVSpanTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.WaterImage.Sp2 = common.IntValue(WaterImageVSpanTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(WaterImageVSpanTb)
	p.Add(row)

}

//背景音乐
func buildFiveteenBgm(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	BgmCb = gwu.NewCheckBox("背景音乐")
	BgmCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AddBgm.Switch = BgmCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(BgmCb)
	//row.AddHSpace(18)
	//descLabel := gwu.NewLabel("是否覆盖:")
	//row.Add(descLabel)
	//
	//BgmLb = gwu.NewListBox([]string{"未选择", "覆盖", "保留",})
	//BgmLb.AddEHandlerFunc(func(e gwu.Event) {
	//
	//	common.AppConfig.AddBgm.Keep = BgmLb.SelectedIdx()
	//
	//}, gwu.ETypeChange)
	//row.Add(BgmLb)

	row.AddHSpace(18)
	descLabel := gwu.NewLabel("将背景音乐放入到bgm文件夹,随机添加")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)
}

//片头片尾
func buildSixteenHeadEnd(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	filmHeadCb = gwu.NewCheckBox("片头")
	filmHeadCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.FilmHead.Switch = filmHeadCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row.Add(filmHeadCb)

	row.AddHSpace(50)
	films, keys := common.LoadFilms()

	arr := []string{"未选择"}
	arr = append(arr, keys...)
	filmHeadLb = gwu.NewListBox(arr)
	filmHeadLb.AddEHandlerFunc(func(e gwu.Event) {
		if filmHeadLb.SelectedIdx() > 0 {
			common.AppConfig.FilmHead.Path = films[filmHeadLb.SelectedValue()]
		}

	}, gwu.ETypeChange)
	row.Add(filmHeadLb)

	row.AddHSpace(20)
	descLabel := gwu.NewLabel("添加片头")
	descLabel.Style().SetColor(GRAY)
	row.Add(descLabel)

	p.Add(row)

	//
	row2 := gwu.NewHorizontalPanel()
	filmFootCb = gwu.NewCheckBox("片尾")
	filmFootCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.FilmFoot.Switch = filmFootCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)
	row2.Add(filmFootCb)

	row2.AddHSpace(50)
	filmFootLb = gwu.NewListBox(arr)
	filmFootLb.AddEHandlerFunc(func(e gwu.Event) {
		if filmFootLb.SelectedIdx() > 0 {
			common.AppConfig.FilmFoot.Path = films[filmFootLb.SelectedValue()]
		}
	}, gwu.ETypeChange)
	row2.Add(filmFootLb)

	row2.AddHSpace(20)
	descLabel2 := gwu.NewLabel("添加片尾")
	descLabel2.Style().SetColor(GRAY)
	row2.Add(descLabel2)

	p.Add(row2)
}

func buildBottomBtn(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	row.AddHSpace(300)
	startBtn := gwu.NewButton("开始处理")
	startBtn.Style().SetFontSize("18")
	startBtn.Style().SetColor(gwu.ClrBlue)

	startBtn.AddEHandlerFunc(func(e gwu.Event) {
		if account.AppAccount.IsActive() == false {
			titleLabel.SetText("请先在激活页面激活软件")
			titleLabel.Style().SetColor("red")
			e.MarkDirty(titleLabel)
			return
		}

		if checkParam() == false {
			titleLabel.SetText("请先在参数页面填写参数")
			titleLabel.Style().SetColor("red")
			e.MarkDirty(titleLabel)
			return
		}

		go app.Engine.DoFactory()

	}, gwu.ETypeClick)

	row.Add(startBtn)
	p.Add(row)

}

func checkParam() bool  {
	if app.Engine.AliYunOss.Endpoint == "" || app.Engine.AliYunOss.BucketName == "" || app.Engine.AliYunOss.BucketDomain == ""  {
		return false
	}

	if app.Engine.AliYunCloud.AccessKeyId == "" || app.Engine.AliYunCloud.AccessKeySecret == "" || app.Engine.AliYunCloud.AppKey == ""  {
		return false
	}
	return true
}