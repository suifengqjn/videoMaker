package app

import (
	"fmt"
	"myTool/ffmpeg"
	"myTool/mylog"
	"path/filepath"
)

//生成字幕文件
func (a *App) createSrtWithVideo(videoPath string, suf string) {

	if !ffmpeg.IsVideo(videoPath) {
		return
	}

	// 1. 提取音频文件
	bjmPath := ffmpeg.ExtractBgm(a.FCmd, videoPath)

	// 2. 上传到oss
	remoteUrl, err := a.AliYunOss.UploadObject(bjmPath, filepath.Base(bjmPath))
	if err != nil {
		mylog.LogDebug("文件上传失败", err)
	}

	// 3.阿里云录音文件识别
	AudioResult := a.AliYunCloud.AliYunAudioRecognition(remoteUrl, a.IntelligentBlock)

	// 4.生成字幕文件
	a.AliYunCloud.AliyunAudioResultMakeSubtitleFile(videoPath, AudioResult, suf)

}

func (a *App) createSrtWithAudio(audioPath string, suf string) {

	fmt.Println("进行字幕生成...")
	// 2. 上传到oss
	remoteUrl, err := a.AliYunOss.UploadObject(audioPath, filepath.Base(audioPath))
	if err != nil {
		mylog.LogDebug("文件上传失败", err)
		return
	}

	// 3.阿里云录音文件识别
	AudioResult := a.AliYunCloud.AliYunAudioRecognition(remoteUrl, a.IntelligentBlock)

	// 4.生成字幕文件
	a.AliYunCloud.AliyunAudioResultMakeSubtitleFile(audioPath, AudioResult, suf)

}

func (a *App)createSrtWithUrl(videoPath,url string, suf string)  {
	// 3.阿里云录音文件识别
	AudioResult := a.AliYunCloud.AliYunAudioRecognition(url, a.IntelligentBlock)

	// 4.生成字幕文件
	a.AliYunCloud.AliyunAudioResultMakeSubtitleFile(videoPath, AudioResult, suf)
}
