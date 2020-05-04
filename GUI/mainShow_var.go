package GUI

import "github.com/icza/gowut/gwu"


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
	ClearWaterCb gwu.CheckBox
	ClearWaterTbx gwu.TextBox
	ClearWaterTby gwu.TextBox
	ClearWaterTbw gwu.TextBox
	ClearWaterTbh gwu.TextBox
)
// 7. 去水印2
//var (
//	ClearWater2Cb gwu.CheckBox
//	ClearWater2Tbx gwu.TextBox
//	ClearWater2Tby gwu.TextBox
//	ClearWater2Tbw gwu.TextBox
//	ClearWater2Tbh gwu.TextBox
//)


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
	DubCB gwu.CheckBox
	SpeecherLb gwu.ListBox //播音人
	VolumeTB  gwu.TextBox //音量
	SpeechRate gwu.TextBox // 语速
	PitchRate gwu.TextBox //语调
	BreakTimeTb gwu.TextBox //间隔时间
)

// 字幕背景
var (
	CoverBgCb gwu.CheckBox
	CoverBgMarginVTb gwu.TextBox
	CoverBjColorTb gwu.TextBox
	CoverBjAlphaTb gwu.TextBox
	CoverHTb gwu.TextBox
)

//字幕属性
var (
	SubTitleCb gwu.CheckBox
	SubTitleFontSizeTb gwu.TextBox
	SubTitleFontColorTb gwu.TextBox
	SubTitleFontBjColorTb gwu.TextBox
	SubTitleFontBjAlphaTb gwu.TextBox
)

// plot
var (
	AIPlotCb gwu.CheckBox
	AIPlotRangeStartTb gwu.TextBox
	AIPlotRangeEndTb gwu.TextBox
	AIPlotDurationTb gwu.TextBox

)

// 12. 文字水印
var (
	WaterTextCb gwu.CheckBox
	WaterTextTb gwu.TextBox
	WaterTextSizeTb gwu.TextBox
	WaterTextColorLb gwu.ListBox
	WaterTextFontLb gwu.ListBox
	WaterTextStyleLb gwu.ListBox
	WaterTextHSpanTb gwu.TextBox
	WaterTextVSpanTb gwu.TextBox

)
// 13. 滚动文字
var (
	RunWaterTextCb gwu.CheckBox
	RunWaterTextTb gwu.TextBox //内容
	RunWaterTextSizeTb gwu.TextBox
	RunWaterTextColorLb gwu.ListBox
	RunWaterTextFontLb gwu.ListBox
	RunWaterTextStyleLb gwu.ListBox
	RunWaterTextDirectionLb gwu.ListBox
	RunWaterTextVSpanTb gwu.TextBox

)
// 14. 图片水印
var (
	WaterImageCb gwu.CheckBox
	WaterImageLb gwu.ListBox
	WaterImageStyleLb gwu.ListBox
	WaterImageHSpanTb gwu.TextBox
	WaterImageVSpanTb gwu.TextBox
)

// 16. 背景音乐
var (
	BgmCb gwu.CheckBox
	BgmFrontVolumeTb gwu.TextBox // 前景音量
	BgmBackVolumeTb gwu.TextBox // 背景音量
	BgmCoverCb gwu.CheckBox //是否覆盖
)

var (
	filmHeadCb  gwu.CheckBox
	filmHeadLb gwu.ListBox

	filmFootCb  gwu.CheckBox
	filmFootLb gwu.ListBox
)
