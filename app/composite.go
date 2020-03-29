package app

import (
	"fmt"
	"myTool/common"
	"myTool/ffmpeg"
	"myTool/file"
	"myTool/imageSplice"
	"strings"
)

/*
style
1. 单视频加文案
	1.1 文案转语音
	1.2 语音转字幕
	1.3 合成
2. 多视频加文案
	2.1 文案转语音
	2.2 语音转字幕
	2.3 根据语音长度合成相同长度的视频
	2.4 合成
3. 图片加文案
	3.1 文案转语音
	3.2 语音转字幕
	3.3 根据语音长度合成相同长度的视频
	3.4 合成
4. 仅视频
	4.1 提取字幕，然后修改
	4.2 字幕转语音
	4.3 语音转字幕
	4.4 合成
*/

func (a *App) composite(dir string) string {
	style := a.Composite.Style
	video := ""
	if style > 0 {
		fmt.Println("进行视频原创处理", dir)
	}
	switch style {
		case 1:
			video = a.compositeStyle1(dir)
		case 2:
			video = a.compositeStyle2(dir)
		case 3:
			video = a.compositeStyle3(dir)
		case 4:
			video = a.compositeStyle1(dir)
		default:
		fmt.Println("没有选择视频合成模式")
	}
	if video == "" {
		return ""
	}
	output := fmt.Sprintf("%v/%v_output.%v", dir, file.GetFileBaseName(video), file.GetFileSuf(video))
	file.MoveFile(video, output)

	return output
}

func (a *App) compositeVideo(videoPath, dir string) string {

	fmt.Println("进行视频合成...")
	//添加音频
	audio := getVoicePath(dir)
	info, err := ffmpeg.GetVideoInfo(a.FCmd, videoPath)
	if err != nil {
		return videoPath
	}
	resultPath := info.AddBgmShortest(a.FCmd, videoPath, audio, true)

	//添加字幕
	return a.AddSubTitle(resultPath, dir)

}

func (a *App) compositeStyle1(dir string) string {

	// 1. 文字转语音
	if getVoicePath(dir) == "" {
		a.createVoice(dir)
	}

	// 2. 语音转字幕
	if getSrtPath(dir) == "" {
		audioPath := getVoicePath(dir)
		if len(audioPath) == 0 {
			return ""
		}
		a.createSrtWithAudio(audioPath)
	}

	//只有一个视频
	videoPath := getVideoPath(dir)

	//3.
	return a.compositeVideo(videoPath, dir)

}

func (a *App) compositeStyle2(dir string) string {
	// 1. 文字转语音
	if getVoicePath(dir) == "" {
		a.createVoice(dir)
	}

	// 2. 语音转字幕
	audioPath := ""
	if getSrtPath(dir) == "" {
		audioPath = getVoicePath(dir)
		if len(audioPath) == 0 {
			return ""
		}
		a.createSrtWithAudio(audioPath)
	} else {
		audioPath = getVoicePath(dir)
	}

	//音频长度
	info, err := ffmpeg.GetVideoInfo(a.FCmd, audioPath)
	if err != nil {
		return ""
	}

	videos := getVideos(dir)
	if len(videos) == 0 {
		fmt.Println("当前文件夹没有视频", dir)
		return ""
	}
	output := dir + "/" + common.GetRandomString(6) + "." + file.GetFileSuf(videos[0])
	//视频随机合并
	common.Shuffle(videos)
	output, _ = ffmpeg.MergeMultiVideoByResolution(a.FCmd, videos, output, 0, 0)

	videoInfo, err := ffmpeg.GetVideoInfo(a.FCmd, output)
	if err != nil {
		return ""
	}

	//如果视频的长度大于音频 ,视频后面部分剪掉
	// 比如 视频10s  配音6秒，则视频留 6+2 = 8s
	if videoInfo.Duration > info.Duration {
		dura := videoInfo.Duration - info.Duration + 2
		output = videoInfo.CutBack(a.FCmd, output, dura)
	}
	//3.
	return a.compositeVideo(output, dir)

}

// 图片加文案
func (a *App) compositeStyle3(dir string) string {

	// 1. 文字转语音
	if getVoicePath(dir) == "" {
		a.createVoice(dir)
	}

	// 2. 语音转字幕
	if getSrtPath(dir) == "" {
		audioPath := getVoicePath(dir)
		if len(audioPath) == 0 {
			return ""
		}
		a.createSrtWithAudio(audioPath)
	}

	info, err := ffmpeg.GetVideoInfo(a.FCmd, getVoicePath(dir))
	if err != nil {
		return ""
	}

	//图片合成视频
	images := getImages(dir)
	common.Shuffle(images)
	videoPath := ffmpeg.CreateVideosByImages(a.FCmd, images, 2, 4, info.Duration + 2)

	//3. 合成
	return a.compositeVideo(videoPath, dir)

}

//添加字幕
func (a *App) AddSubTitle(videoPath, dir string) string {

	fmt.Println("添加字幕中...")
	info, err := ffmpeg.GetVideoInfo(a.FCmd, videoPath)
	if err != nil {
		return videoPath
	}
	// 覆盖原有字幕
	if a.Subtitles.CoverBj {
		videoPath = a.coverOldSubTitle(info, info.W, a.Subtitles.CoverH)

		info, err = ffmpeg.GetVideoInfo(a.FCmd, videoPath)
		if err != nil {
			fmt.Println("字幕覆盖失败")
		}
	}

	// 获取字幕文件
	srt := getSrtPath(dir)
	if strings.HasSuffix(srt, "srt") == false {
		fmt.Println("字幕文件不存在", srt)
	}

	fontColor := a.Subtitles.FontColor
	if len(fontColor) < 6 {
		fontColor = "000000"
	}
	alignment := 2
	fontSize := a.Subtitles.FontSize
	if fontSize <= 0 {
		fontSize = 20
	}
	marginV := a.Subtitles.MarginV
	bjColor := a.Subtitles.BjColor
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

func (a *App) coverOldSubTitle(info *ffmpeg.VideoInfo, videoW int, h int) string {

	// 计算黑边尺寸
	lineHeight := h
	imagePath := ffmpeg.Snip(a.FCmd, info.VideoPath, "20", "1")
	if h <= 0 {
		_, lineHeight, _ = imageSplice.BottomLineHeight(imagePath)
		if lineHeight == 0 {
			lineHeight = 100
		}
	}

	//构造黑边
	imageLine := ffmpeg.MakeRandExportPath("jpg")
	err := imageSplice.CreateImage(videoW, lineHeight, "#000000", imageLine)

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
