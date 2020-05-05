package GUI

import (
	"github.com/icza/gowut/gwu"
	"myProject/videoMaker/account"
	"myProject/videoMaker/common"
	"strings"
)

var GRAY = "#808080"
var titleLabel gwu.Label
var defaultShowString = "原创参数配置, 首次使用建议先看教程"

func buildMainShowUI(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	buildUI(p)

	return p
}

func buildUI(p gwu.Panel) {

	topLine := gwu.NewHorizontalPanel()

	titleLabel = gwu.NewLabel(defaultShowString)
	titleLabel.Style().SetColor(gwu.ClrBlue)
	topLine.Add(titleLabel)

	topLine.AddHSpace(30)
	link0 := gwu.NewLink("播音人", "https://github.com/suifengqjn/videoMaker/blob/master/voicer.md")
	topLine.Add(link0)

	topLine.AddHSpace(30)
	link := gwu.NewLink("颜色值选取", "http://cha.buyiju.com/tool/color.html")
	topLine.Add(link)

	p.Add(topLine)

	//
	//p.AddVSpace(10)
	//buildSevenClearWater2(p)

	//-------//
	p.AddVSpace(10)
	buildSrtUI(p)

	p.AddVSpace(10)
	buildComposite(p)

	p.AddVSpace(10)
	buildDub(p)

	p.AddVSpace(10)
	buildSubTitleBg(p)

	p.AddVSpace(10)
	buildSubTitle(p)

	p.AddVSpace(10)
	buildAIPlot(p)
	//-------//

	p.AddVSpace(10)
	buildFourCutFront(p)

	p.AddVSpace(10)
	buildSevenClearWater(p)

	p.AddVSpace(10)
	buildTwelveTextWater(p)

	//p.AddVSpace(10)
	//buildTwelveTextWater2(p)

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

	}, gwu.ETypeClick, gwu.ETypeStateChange)

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

	row.AddHSpace(20)
	cutBackCb = gwu.NewCheckBox("剪去片尾")
	cutBackCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.CutBack.Switch = cutBackCb.State()

	}, gwu.ETypeClick, gwu.ETypeStateChange)

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
	descLabel = gwu.NewLabel("如: 7 ,剪去片尾7秒")
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
	}, gwu.ETypeClick, gwu.ETypeStateChange)

	row.Add(SrtCb)
	p.Add(row)
}

func buildComposite(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	CompositeCb = gwu.NewCheckBox("视频合成")
	CompositeCb.Style().SetColor("#FF6633")
	CompositeCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.CompositeStyle.Switch = CompositeCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
	row.Add(CompositeCb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("模式:"))
	CompositeLb = gwu.NewListBox(common.CompleteStyle)
	CompositeLb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.CompositeStyle.Style = CompositeLb.SelectedIdx()
	}, gwu.ETypeChange)
	row.Add(CompositeLb)

	row.AddHSpace(10)

	p.Add(row)

	//p.AddVSpace(10)
	//row2 := gwu.NewHorizontalPanel()
	//row2.AddHSpace(100)
	//desc := gwu.NewLabel("")
	//desc.Style().SetColor(GRAY)
	//row2.Add(desc)

}

//配音
func buildDub(p gwu.Panel) {

	row := gwu.NewHorizontalPanel()

	DubCB = gwu.NewCheckBox("配音设置")
	DubCB.Style().SetColor("#FF6633")
	DubCB.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Dub.Switch = DubCB.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
	row.Add(DubCB)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("配音员:"))
	SpeecherLb = gwu.NewListBox(common.Voiceover)
	common.AppConfig.AppConf.Dub.Voice = "Xiaoyun"
	SpeecherLb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Dub.Voice = common.VoiceoverMap[SpeecherLb.SelectedValue()]

	}, gwu.ETypeChange)
	row.Add(SpeecherLb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("音量:"))
	VolumeTB = gwu.NewTextBox("50")
	VolumeTB.SetMaxLength(3)
	VolumeTB.Style().SetWidthPx(50)
	VolumeTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	common.AppConfig.AppConf.Dub.Volume = 50
	VolumeTB.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AppConf.Dub.Volume = common.IntValue(VolumeTB.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(VolumeTB)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("语速:"))
	SpeechRate = gwu.NewTextBox("0")
	SpeechRate.SetMaxLength(3)
	SpeechRate.Style().SetWidthPx(50)
	SpeechRate.AddSyncOnETypes(gwu.ETypeKeyUp)
	SpeechRate.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Dub.SpeechRate = common.IntValue(SpeechRate.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SpeechRate)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("语调:"))
	PitchRate = gwu.NewTextBox("0")
	PitchRate.SetMaxLength(3)
	PitchRate.Style().SetWidthPx(50)
	PitchRate.AddSyncOnETypes(gwu.ETypeKeyUp)
	PitchRate.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Dub.PitchRate = common.IntValue(PitchRate.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(PitchRate)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("停顿:"))
	BreakTimeTb = gwu.NewTextBox("")
	BreakTimeTb.SetMaxLength(4)
	BreakTimeTb.Style().SetWidthPx(50)
	BreakTimeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	BreakTimeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Dub.BreakTime = common.IntValue(BreakTimeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(BreakTimeTb)

	row.AddHSpace(10)
	breakDesc := gwu.NewLabel("音量(0~100)，默认50. 语速和语调(-500~500)，默认是0, 停顿单位毫秒")
	breakDesc.Style().SetColor(GRAY)
	row.Add(breakDesc)

	p.Add(row)
}

//字幕遮盖条
func buildSubTitleBg(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	CoverBgCb = gwu.NewCheckBox("字幕遮盖")
	CoverBgCb.Style().SetColor("#FF6633")
	CoverBgCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.Switch = CoverBgCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
	row.Add(CoverBgCb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("底部距离:"))
	CoverBgMarginVTb = gwu.NewTextBox("10")
	CoverBgMarginVTb.SetMaxLength(3)
	CoverBgMarginVTb.Style().SetWidthPx(50)
	CoverBgMarginVTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	CoverBgMarginVTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.SubtitleBack.CoverB = common.IntValue(CoverBgMarginVTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.AppConf.SubtitleBack.CoverB = 10
	row.Add(CoverBgMarginVTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("遮盖条颜色:"))
	CoverBjColorTb = gwu.NewTextBox("00ff00")
	CoverBjColorTb.SetMaxLength(7)
	CoverBjColorTb.Style().SetWidthPx(50)
	CoverBjColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	CoverBjColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.SubtitleBack.BjColor = CoverBjColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.AppConf.SubtitleBack.BjColor = "00ff00"
	row.Add(CoverBjColorTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("遮盖高度:"))
	CoverHTb = gwu.NewTextBox("30")
	CoverHTb.SetMaxLength(3)
	CoverHTb.Style().SetWidthPx(50)
	CoverHTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	CoverHTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.SubtitleBack.CoverH = common.IntValue(CoverHTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.AppConf.SubtitleBack.CoverH = 30
	row.Add(CoverHTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("遮盖条透明度:"))
	CoverBjAlphaTb = gwu.NewTextBox("10")
	CoverBjAlphaTb.SetMaxLength(3)
	CoverBjAlphaTb.Style().SetWidthPx(50)
	CoverBjAlphaTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	CoverBjAlphaTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.SubtitleBack.BjAlpha = common.IntValue(CoverBjAlphaTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.AppConf.SubtitleBack.BjAlpha = 10
	row.Add(CoverBjAlphaTb)

	row.AddHSpace(10)
	desc := gwu.NewLabel("透明度 0-10, 0表示完全透明, 10不透明")
	desc.Style().SetColor(GRAY)
	row.Add(desc)

	p.Add(row)
}

// 字幕属性
func buildSubTitle(p gwu.Panel) {

	row := gwu.NewHorizontalPanel()
	SubTitleCb = gwu.NewCheckBox("字幕属性")
	SubTitleCb.Style().SetColor("#FF6633")
	SubTitleCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.Switch = SubTitleCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
	row.Add(SubTitleCb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字体大小:"))
	SubTitleFontSizeTb = gwu.NewTextBox("16")
	SubTitleFontSizeTb.SetMaxLength(3)
	SubTitleFontSizeTb.Style().SetWidthPx(50)
	SubTitleFontSizeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontSizeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.FontSize = common.IntValue(SubTitleFontSizeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.Subtitles.FontSize = 16
	row.Add(SubTitleFontSizeTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字体颜色:"))
	SubTitleFontColorTb = gwu.NewTextBox("000000")
	SubTitleFontColorTb.SetMaxLength(7)
	SubTitleFontColorTb.Style().SetWidthPx(50)
	SubTitleFontColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.Subtitles.FontColor = SubTitleFontColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	common.AppConfig.Subtitles.FontColor = "000000"
	row.Add(SubTitleFontColorTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字幕背景颜色:"))
	SubTitleFontBjColorTb = gwu.NewTextBox("")
	SubTitleFontBjColorTb.SetMaxLength(7)
	SubTitleFontBjColorTb.Style().SetWidthPx(50)
	SubTitleFontBjColorTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontBjColorTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Subtitles.BjColor = SubTitleFontBjColorTb.Text()
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SubTitleFontBjColorTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("字幕条背景透明度:"))
	SubTitleFontBjAlphaTb = gwu.NewTextBox("")
	SubTitleFontBjAlphaTb.SetMaxLength(7)
	SubTitleFontBjAlphaTb.Style().SetWidthPx(50)
	SubTitleFontBjAlphaTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	SubTitleFontBjAlphaTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.Subtitles.BjAlpha = common.IntValue(SubTitleFontBjAlphaTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(SubTitleFontBjAlphaTb)

	row.AddHSpace(10)
	desc := gwu.NewLabel("透明度 0-10, 0表示完全透明, 10不透明")
	desc.Style().SetColor(GRAY)
	row.Add(desc)

	p.Add(row)

}


// AI剪辑策略
func buildAIPlot(p gwu.Panel) {

	row := gwu.NewHorizontalPanel()
	AIPlotCb = gwu.NewCheckBox("剪辑策略")
	AIPlotCb.Style().SetColor("#FF6633")
	AIPlotCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AIPlot.Switch = AIPlotCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
	row.Add(AIPlotCb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("起始范围:"))
	AIPlotRangeStartTb = gwu.NewTextBox("")
	AIPlotRangeStartTb.SetMaxLength(3)
	AIPlotRangeStartTb.Style().SetWidthPx(50)
	AIPlotRangeStartTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AIPlotRangeStartTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AIPlot.RangeStart = common.IntValue(AIPlotRangeStartTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AIPlotRangeStartTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("终止范围:"))
	AIPlotRangeEndTb = gwu.NewTextBox("")
	AIPlotRangeEndTb.SetMaxLength(7)
	AIPlotRangeEndTb.Style().SetWidthPx(50)
	AIPlotRangeEndTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AIPlotRangeEndTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AIPlot.RangeEnd = common.IntValue(AIPlotRangeEndTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AIPlotRangeEndTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("视频时长:"))
	AIPlotDurationTb = gwu.NewTextBox("")
	AIPlotDurationTb.SetMaxLength(3)
	AIPlotDurationTb.Style().SetWidthPx(50)
	AIPlotDurationTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	AIPlotDurationTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.AIPlot.Duration = common.IntValue(AIPlotDurationTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(AIPlotDurationTb)

	row.AddHSpace(10)
	desc := gwu.NewLabel("如果限定时长，多出部分会被剪掉")
	desc.Style().SetColor(GRAY)
	row.Add(desc)

	p.Add(row)

}

//去除水印
func buildSevenClearWater(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()
	ClearWaterCb = gwu.NewCheckBox("去除水印")
	ClearWaterCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.ClearWater.Switch = ClearWaterCb.State()

	}, gwu.ETypeClick,gwu.ETypeStateChange)

	row.Add(ClearWaterCb)
	row.AddHSpace(18)

	xl := gwu.NewLabel("x:")
	row.Add(xl)
	ClearWaterTbx= gwu.NewTextBox("")
	ClearWaterTbx.SetMaxLength(5)
	ClearWaterTbx.Style().SetWidthPx(50)
	ClearWaterTbx.AddSyncOnETypes(gwu.ETypeKeyUp)
	ClearWaterTbx.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AppConf.ClearWater.X = common.IntValue(ClearWaterTbx.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(ClearWaterTbx)

	row.AddHSpace(5)
	yl := gwu.NewLabel("y:")
	row.Add(yl)
	ClearWaterTby= gwu.NewTextBox("")
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
	ClearWaterTbw= gwu.NewTextBox("")
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

//文字水印
func buildTwelveTextWater(p gwu.Panel) {
	row1 := gwu.NewHorizontalPanel()
	WaterTextCb = gwu.NewCheckBox("文字水印")
	WaterTextCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.WaterText.Switch = WaterTextCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)

	row1.Add(WaterTextCb)

	row1.AddHSpace(18)
	content := gwu.NewLabel("内容:")
	row1.Add(content)
	WaterTextTb = gwu.NewTextBox("")
	WaterTextTb.SetToolTip("123123")
	WaterTextTb.SetMaxLength(500)
	WaterTextTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	WaterTextTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AppConf.WaterText.Content = WaterTextTb.Text()
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


//图片水印
func buildThirteenImageWater(p gwu.Panel) {
	row := gwu.NewHorizontalPanel()

	WaterImageCb = gwu.NewCheckBox("图片水印")
	WaterImageCb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.WaterImage.Switch = WaterImageCb.State()
	}, gwu.ETypeClick, gwu.ETypeStateChange)
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

	}, gwu.ETypeClick, gwu.ETypeStateChange)

	row.Add(BgmCb)


	row.AddHSpace(10)
	row.Add(gwu.NewLabel("前景音量:"))
	BgmFrontVolumeTb = gwu.NewTextBox("")
	BgmFrontVolumeTb.SetMaxLength(3)
	BgmFrontVolumeTb.Style().SetWidthPx(50)
	BgmFrontVolumeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	BgmFrontVolumeTb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AddBgm.FrontVolume = common.FloatValue(BgmFrontVolumeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(BgmFrontVolumeTb)

	row.AddHSpace(10)
	row.Add(gwu.NewLabel("背景音量:"))
	BgmBackVolumeTb = gwu.NewTextBox("")
	BgmBackVolumeTb.SetMaxLength(7)
	BgmBackVolumeTb.Style().SetWidthPx(50)
	BgmBackVolumeTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	BgmBackVolumeTb.AddEHandlerFunc(func(e gwu.Event) {
		common.AppConfig.AddBgm.BackVolume = common.FloatValue(BgmBackVolumeTb.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(BgmBackVolumeTb)

	row.AddHSpace(10)
	BgmCoverCb = gwu.NewCheckBox("是否覆盖")
	BgmCoverCb.AddEHandlerFunc(func(e gwu.Event) {

		common.AppConfig.AddBgm.Cover = BgmCoverCb.State()

	}, gwu.ETypeClick, gwu.ETypeStateChange)

	row.Add(BgmCoverCb)

	row.AddHSpace(20)
	descLabel := gwu.NewLabel("音量范围0-3，1为原始音量，小于1音量变小， 大于1音量变大")
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

	}, gwu.ETypeClick, gwu.ETypeStateChange)
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

	}, gwu.ETypeClick, gwu.ETypeStateChange)
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

		go common.MakerEngine.DoMaker()

	}, gwu.ETypeClick)

	row.Add(startBtn)
	p.Add(row)

}

func checkParam() bool {
	if common.MakerEngine.MakerCli.AliYunOss.Endpoint == "" || common.MakerEngine.MakerCli.AliYunOss.BucketName == "" || common.MakerEngine.MakerCli.AliYunOss.BucketDomain == "" {
		return false
	}

	if common.MakerEngine.MakerCli.AliYunCloud.AccessKeyId == "" || common.MakerEngine.MakerCli.AliYunCloud.AccessKeySecret == "" || common.MakerEngine.MakerCli.AliYunCloud.AppKey == "" {
		return false
	}
	return true
}
