package app

import (
	"fmt"
	"myTool/ffmpeg"
	"myTool/file"
	"myTool/imageSplice"
	"os"
	"strings"
)

var IsDealing bool
func (a *App) DoFactory() {

	if IsDealing {
		fmt.Println("软件正在处理中，请处理完再开始新的处理")
		return
	}

	videos := a.getVideos()
	if len(videos) == 0 {
		fmt.Println("没有视频需要处理")
	}
	IsDealing = true
	for _, f := range videos {
		a.editVideo(f)
	}

	fmt.Printf("视频处理结束, 一共处理 %v 个视频", len(videos))
	IsDealing = false

}

func (a *App) editVideo(f string) bool {

	//预处理
	f = a.prepareEdit(f)

	//原创处理
	f = a.originalEdit(f)

	// 后处理
	f = a.postEdit(f)


	return true

}

func (a *App) prepareEdit(f string) string {

	fCmd := a.FCmd
	tempf := f
	if a.CutFront.Switch && a.CutBack.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			if a.CutFront.Value > 0 && a.CutBack.Value > 0 {
				f = info.CutFrontAndBack(fCmd, f, a.CutFront.Value, a.CutBack.Value)
			}
		}

	} else {
		//4. cut front
		if a.CutFront.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err == nil && a.CutFront.Value > 0 {
				f = info.CutFront(fCmd, f, a.CutFront.Value)
			}

		}

		//5. cut back
		if a.CutBack.Switch {
			info, err := ffmpeg.GetVideoInfo(fCmd, f)
			if err == nil && a.CutBack.Value > 0 {
				f = info.CutBack(fCmd, f, a.CutBack.Value)
			}

		}
	}

	//去除水印
	if a.ClearWater.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			f = info.ClearWater(fCmd, f, a.ClearWater.X, a.ClearWater.Y, a.ClearWater.W, a.ClearWater.H)
		}

	}

	if a.ClearWater1.Switch {
		info, err := ffmpeg.GetVideoInfo(fCmd, f)
		if err == nil {
			f = info.ClearWater(fCmd, f, a.ClearWater1.X, a.ClearWater1.Y, a.ClearWater1.W, a.ClearWater1.H)
		}

	}

	//覆盖文件
	if tempf != f {
		os.Remove(tempf)
		file.MoveFile(f, tempf)
	}
	return tempf
}

func (a *App) originalEdit(f string) string {

	tempf := f
	// 1. 生成字幕文件
	if a.ExtractSubtitles.Switch {
		a.createSrt(f)
	}

	// 2. 字幕或者文本转语音
	if a.TextToVoice.Switch {
		url ,_ :=  a.createVoice(f)
		// 3. 通过新的语音再次生成字幕文件
		a.createSrtWithUrl(f, url)
	}

	if a.Composite.Switch {
		f = a.composite(f)
	}

	// 覆盖
	if tempf != f {
		os.Remove(tempf)
		file.MoveFile(f, tempf)
	}

	return tempf
}

func (a *App)composite(videoPath string) string {
	style := a.Composite.Style
	if style == 1 {  // 配音加字幕
		videoPath = a.coverDubbing(getVoicePath(videoPath), videoPath)
		videoPath = a.AddSubTitle(videoPath)
	} else if style == 2 {  //仅配音
		videoPath = a.coverDubbing(getVoicePath(videoPath), videoPath)
	} else if style == 3 {  //仅字幕
		videoPath = a.AddSubTitle(videoPath)
	}

	return videoPath

}

//覆盖配音
func (a *App)coverDubbing(voicePath, videoPath string)string  {
	info, err := ffmpeg.GetVideoInfo(a.FCmd,videoPath)
	if err != nil {
		return videoPath
	}
	return info.AddBgm(a.FCmd,videoPath,voicePath,true)
}

//添加字幕
func (a *App)AddSubTitle(videoPath string) string  {

	info, err := ffmpeg.GetVideoInfo(a.FCmd,videoPath)
	if err != nil {
		return videoPath
	}
	// 覆盖原有字幕
	if a.Subtitles.CoverBj {
		videoPath = a.coverOldSubTitle(info, info.W,a.Subtitles.CoverH)
	}

	// 获取字幕文件
	srt := getSrtPath(videoPath)
	if strings.HasSuffix(srt,"srt") == false {
		fmt.Println("字幕文件不存在", srt)
	}

	fontColor := a.Subtitles.FontColor
	if len(fontColor) < 6 {
		fontColor = "ffffff"
	}
	alignment := 2
	fontSize := a.Subtitles.FontSize
	if fontSize <= 0 {
		fontSize = 20
	}
	marginV := a.Subtitles.MarginV
	bjColor :=  a.Subtitles.BjColor
	bjColorAlpha := a.Subtitles.BjAlpha
	videoPath = info.AddSubTitle(a.FCmd,
		srt,
		fontColor,
		alignment,
		fontSize,
		marginV,
		0,
		"",
		bjColor,
		bjColorAlpha,
		)
	return videoPath
}


func (a *App)coverOldSubTitle(info *ffmpeg.VideoInfo,videoW int,h int) string {

	// 计算黑边尺寸
	lineHeight := h
	imagePath := ffmpeg.Snip(a.FCmd,info.VideoPath,"20","1")
	if h <= 0 {
		_,lineHeight,_ = imageSplice.BottomLineHeight(imagePath)
		if lineHeight == 0{
			lineHeight = 100
		}
	}

	//构造黑边
	imageLine := ffmpeg.MakeRandExportPath("jpg")
	err := imageSplice.CreateImage(videoW, lineHeight,"#000000",imageLine)

	if err != nil {
		return ""
	}

	//将黑边添加到底部
	f := info.AddTextWaterImageWithStyle(
		a.FCmd,
		info.VideoPath,
		imageLine,
		4,
		0,
		0,
	)

	return f
}

func (a *App)postEdit(voicePath string) string  {

	return ""
}