package GUI

import (
	"github.com/icza/gowut/gwu"
	cm "myProject/videoCli/common"
	"myProject/videoMaker/common"
)

// 4. 片头
var (
	cutHeadCb gwu.CheckBox
	cutHeadTb gwu.TextBox
)

// 5. 片尾
var (
	cutBackCb gwu.CheckBox
	cutBackTb gwu.TextBox
)

// 7. 去水印1
var (
	ClearWaterCb  gwu.CheckBox
	ClearWaterTbx gwu.TextBox
	ClearWaterTby gwu.TextBox
	ClearWaterTbw gwu.TextBox
	ClearWaterTbh gwu.TextBox
)

// 提取字幕
var (
	SrtCb gwu.CheckBox
)

//合成  4种模式
var (
	CompositeCb gwu.CheckBox
	CompositeLb gwu.ListBox
)

//配音
/*
	Voice string    //发音人，默认是xiaoyun
	Volume int      //音量，范围是0~100，默认50
	SpeechRate int  //语速，范围是-500~500，默认是0
	PitchRate int   //语调，范围是-500~500，默认是0
*/
var (
	DubCB       gwu.CheckBox
	SpeecherLb  gwu.ListBox //播音人
	VolumeTB    gwu.TextBox //音量
	SpeechRate  gwu.TextBox // 语速
	PitchRate   gwu.TextBox //语调
	BreakTimeTb gwu.TextBox //间隔时间
)

// 字幕背景
var (
	CoverBgCb        gwu.CheckBox
	CoverBgMarginVTb gwu.TextBox
	CoverBjColorTb   gwu.TextBox
	CoverBjAlphaTb   gwu.TextBox
	CoverHTb         gwu.TextBox
)

//字幕属性
var (
	SubTitleCb            gwu.CheckBox
	SubTitleFontVTb       gwu.TextBox
	SubTitleFontSizeTb    gwu.TextBox
	SubTitleFontColorTb   gwu.TextBox
	SubTitleFontBjColorTb gwu.TextBox
	SubTitleFontBjAlphaTb gwu.TextBox
)

// plot
var (
	AIPlotCb           gwu.CheckBox
	AIPlotRangeStartTb gwu.TextBox
	AIPlotRangeEndTb   gwu.TextBox
	AIPlotDurationTb   gwu.TextBox
)

// 12. 文字水印
var (
	WaterTextCb      gwu.CheckBox
	WaterTextTb      gwu.TextBox
	WaterTextSizeTb  gwu.TextBox
	WaterTextColorLb gwu.ListBox
	WaterTextFontLb  gwu.ListBox
	WaterTextStyleLb gwu.ListBox
	WaterTextHSpanTb gwu.TextBox
	WaterTextVSpanTb gwu.TextBox
)


// 14. 图片水印
var (
	WaterImageCb      gwu.CheckBox
	WaterImageLb      gwu.ListBox
	WaterImageStyleLb gwu.ListBox
	WaterImageHSpanTb gwu.TextBox
	WaterImageVSpanTb gwu.TextBox
)

// 16. 背景音乐
var (
	BgmCb            gwu.CheckBox
	BgmFrontVolumeTb gwu.TextBox  // 前景音量
	BgmBackVolumeTb  gwu.TextBox  // 背景音量
	BgmCoverCb       gwu.CheckBox //是否覆盖
)

var (
	filmHeadCb gwu.CheckBox
	filmHeadLb gwu.ListBox

	filmFootCb gwu.CheckBox
	filmFootLb gwu.ListBox
)

func fillWithConfig(con *cm.MakerConfig,e gwu.Event) {

	cutHeadCb.SetState(con.CutFront.Switch)
	cutHeadCb.SetText(common.StrValue(con.CutFront.Value))
	e.MarkDirty(cutHeadCb)
	e.MarkDirty(cutHeadTb)

	cutBackCb.SetState(con.CutBack.Switch)
	cutBackTb.SetText(common.StrValue(con.CutBack.Value))
	e.MarkDirty(cutBackCb)
	e.MarkDirty(cutBackTb)

	ClearWaterCb.SetState(con.ClearWater.Switch)
	ClearWaterTbx.SetText(common.StrValue(con.ClearWater.X))
	ClearWaterTby.SetText(common.StrValue(con.ClearWater.Y))
	ClearWaterTbw.SetText(common.StrValue(con.ClearWater.W))
	ClearWaterTbh.SetText(common.StrValue(con.ClearWater.W))
	e.MarkDirty(ClearWaterCb)
	e.MarkDirty(ClearWaterTbx)
	e.MarkDirty(ClearWaterTby)
	e.MarkDirty(ClearWaterTbw)
	e.MarkDirty(ClearWaterTbh)

	SrtCb.SetState(con.ExtractSubtitles.Switch)
	e.MarkDirty(SrtCb)


	CompositeCb.SetState(con.CompositeStyle.Switch)
	CompositeLb.ClearSelected()
	for _, v := range common.CompleteStyleMap {
		if con.CompositeStyle.Style == v {
			CompositeLb.Selected(v)
			break
		}
	}
	e.MarkDirty(CompositeCb)
	e.MarkDirty(CompositeLb)


	//voice
	DubCB.SetState(con.Dub.Switch)
	SpeecherLb.ClearSelected()
	for k, v := range common.VoiceoverMap {
		if con.Dub.Voice == v {
			for index, p := range common.Voiceover {
				if p == k {
					SpeecherLb.Selected(index)
					break
				}
			}
		}
	}
	VolumeTB.SetText(con.Dub.Voice)
	SpeechRate.SetText(common.StrValue(con.Dub.SpeechRate))
	PitchRate.SetText(common.StrValue(con.Dub.PitchRate))
	BreakTimeTb.SetText(common.StrValue(con.Dub.BreakTime))
	e.MarkDirty(DubCB)
	e.MarkDirty(SpeecherLb)
	e.MarkDirty(VolumeTB)
	e.MarkDirty(SpeechRate)
	e.MarkDirty(PitchRate)
	e.MarkDirty(BreakTimeTb)

	// 字幕背景
	CoverBgCb.SetState(con.SubtitleBack.Switch)
	CoverBgMarginVTb.SetText(common.StrValue(con.SubtitleBack.CoverB))
	CoverBjColorTb.SetText(con.SubtitleBack.BjColor)
	CoverBjAlphaTb.SetText(common.StrValue(con.SubtitleBack.BjAlpha))
	CoverHTb.SetText(common.StrValue(con.SubtitleBack.CoverH))
	e.MarkDirty(CoverBgCb)
	e.MarkDirty(CoverBgMarginVTb)
	e.MarkDirty(CoverBjColorTb)
	e.MarkDirty(CoverBjAlphaTb)
	e.MarkDirty(CoverHTb)


	//字幕属性
	SubTitleCb.SetState(con.Subtitles.Switch)
	SubTitleFontVTb.SetText(common.StrValue(con.Subtitles.MarginV))
	SubTitleFontSizeTb.SetText(common.StrValue(con.Subtitles.FontSize))
	SubTitleFontColorTb.SetText(con.Subtitles.FontColor)
	SubTitleFontBjColorTb.SetText(con.Subtitles.BjColor)
	SubTitleFontBjAlphaTb.SetText(common.StrValue(con.Subtitles.BjAlpha))
	e.MarkDirty(SubTitleCb)
	e.MarkDirty(SubTitleFontVTb)
	e.MarkDirty(SubTitleFontSizeTb)
	e.MarkDirty(SubTitleFontColorTb)
	e.MarkDirty(SubTitleFontBjColorTb)
	e.MarkDirty(SubTitleFontBjAlphaTb)


	//plot
	AIPlotCb.SetState(con.AIPlot.Switch)
	AIPlotRangeStartTb.SetText(common.StrValue(con.AIPlot.RangeStart))
	AIPlotRangeEndTb.SetText(common.StrValue(con.AIPlot.RangeEnd))
	AIPlotDurationTb.SetText(common.StrValue(con.AIPlot.Duration))
	e.MarkDirty(AIPlotCb)
	e.MarkDirty(AIPlotRangeStartTb)
	e.MarkDirty(AIPlotRangeEndTb)
	e.MarkDirty(AIPlotDurationTb)


}
