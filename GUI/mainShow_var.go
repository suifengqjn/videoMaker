package GUI

import (
	"fmt"
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

	//提取字幕
	SrtCb.SetState(con.ExtractSubtitles.Switch)
	e.MarkDirty(SrtCb)

	//视频合成
	CompositeCb.SetState(con.CompositeStyle.Switch)
	CompositeLb.ClearSelected()
	for _, v := range common.CompleteStyleMap {
		if con.CompositeStyle.Style == v {
			CompositeLb.SetSelected(v, true)
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
					SpeecherLb.SetSelected(index, true)
					break
				}
			}
		}
	}
	VolumeTB.SetText(common.StrValue(con.Dub.Volume))
	SpeechRate.SetText(common.StrValue(con.Dub.SpeechRate))
	PitchRate.SetText(common.StrValue(con.Dub.PitchRate))
	BreakTimeTb.SetText(common.StrValue(con.Dub.BreakTime))
	e.MarkDirty(DubCB)
	e.MarkDirty(SpeecherLb)
	e.MarkDirty(VolumeTB)
	e.MarkDirty(SpeechRate)
	e.MarkDirty(PitchRate)
	e.MarkDirty(BreakTimeTb)

	// 字幕遮盖
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

	//片头片尾
	cutHeadCb.SetState(con.CutFront.Switch)
	cutHeadTb.SetText(common.StrValue(con.CutFront.Value))
	e.MarkDirty(cutHeadCb)
	e.MarkDirty(cutHeadTb)

	cutBackCb.SetState(con.CutBack.Switch)
	cutBackTb.SetText(common.StrValue(con.CutBack.Value))
	e.MarkDirty(cutBackCb)
	e.MarkDirty(cutBackTb)

	//去除水印
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

	//文字水印
	WaterTextCb.SetState(con.WaterText.Switch)
	WaterTextTb.SetText(con.WaterText.Content)
	WaterTextSizeTb.SetText(common.StrValue(con.WaterText.Size))
	if _, ok := common.TextColorsMap[con.WaterText.Color];ok {
		index := common.TextColorsMap[con.WaterText.Color]
		WaterTextColorLb.SetSelected(index, true)
	} else {
		WaterTextColorLb.ClearSelected()
	}

	fonts, keys := common.LoadFonts()
	WaterTextFontLb.ClearSelected()
	for _, v := range fonts {
		if v == con.WaterText.Path {
			for i:= 0;i<len(keys);i++ {
				WaterTextFontLb.SetSelected(i+1, true)
				break
			}
		}
	}


	if _, ok := common.WaterStyleMap[con.WaterText.Style];ok {
		WaterTextStyleLb.SetSelected(con.WaterText.Style, true)
	} else {
		WaterTextStyleLb.ClearSelected()
	}
	WaterTextHSpanTb.SetText(common.StrValue(con.WaterText.Sp1))
	WaterTextVSpanTb.SetText(common.StrValue(con.WaterText.Sp2))
	e.MarkDirty(WaterTextCb)
	e.MarkDirty(WaterTextTb)
	e.MarkDirty(WaterTextSizeTb)
	e.MarkDirty(WaterTextColorLb)
	e.MarkDirty(WaterTextFontLb)
	e.MarkDirty(WaterTextStyleLb)
	e.MarkDirty(WaterTextHSpanTb)
	e.MarkDirty(WaterTextVSpanTb)


	//图片水印
	images, keys := common.LoadImages()
	WaterImageCb.SetState(con.WaterImage.Switch)
	WaterImageLb.ClearSelected()
	for _, v := range images {
		if v == con.WaterImage.Path {
			for i:= 0;i<len(keys);i++ {
				WaterImageLb.SetSelected(i+1, true)
				break
			}
		}
	}
	if _, ok := common.WaterStyleMap[con.WaterImage.Style];ok {
		WaterImageStyleLb.SetSelected(con.WaterImage.Style, true)
	} else {
		WaterImageStyleLb.ClearSelected()
	}
	WaterImageHSpanTb.SetText(common.StrValue(con.WaterImage.Sp1))
	WaterImageVSpanTb.SetText(common.StrValue(con.WaterImage.Sp2))
	e.MarkDirty(WaterImageCb)
	e.MarkDirty(WaterImageLb)
	e.MarkDirty(WaterImageStyleLb)
	e.MarkDirty(WaterImageHSpanTb)
	e.MarkDirty(WaterImageVSpanTb)

	//背景音乐
	BgmCb.SetState(con.AddBgm.Switch)
	BgmFrontVolumeTb.SetText(fmt.Sprintf("%v",con.AddBgm.FrontVolume))
	BgmBackVolumeTb.SetText(fmt.Sprintf("%v",con.AddBgm.BackVolume))
	BgmCoverCb.SetState(con.AddBgm.Cover)
	e.MarkDirty(BgmCb)
	e.MarkDirty(BgmFrontVolumeTb)
	e.MarkDirty(BgmBackVolumeTb)
	e.MarkDirty(BgmCoverCb)

	//片头片尾
	films, keys := common.LoadFilms()
	filmHeadCb.SetState(con.FilmHead.Switch)
	filmHeadLb.ClearSelected()
	for _, v := range films {
		if v == con.FilmHead.Path {
			for i:= 0;i<len(keys);i++ {
				filmHeadLb.SetSelected(i+1, true)
				break
			}
		}
	}
	e.MarkDirty(filmHeadCb)
	e.MarkDirty(filmHeadLb)

	filmFootCb.SetState(con.FilmFoot.Switch)
	filmFootLb.ClearSelected()
	for _, v := range films {
		if v == con.FilmFoot.Path {
			for i:= 0;i<len(keys);i++ {
				filmFootLb.SetSelected(i+1, true)
				break
			}
		}
	}
	e.MarkDirty(filmFootCb)
	e.MarkDirty(filmFootLb)














}
