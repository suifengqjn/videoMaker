package app

import (
	"myTool/ffmpeg"
	"myTool/mylog"
	"path/filepath"
)

//生成字幕文件
func (a *App) createSrt(videoPath string) {

	if !ffmpeg.IsVideo(videoPath) {
		return
	}

	// 1. 提取音频文件
	bjmPath := ffmpeg.ExtractBgm(a.AppConfig.FCmd, videoPath)

	// 2. 上传到oss
	remoteUrl, err := a.AppConfig.AliYunOss.UploadObject(bjmPath, filepath.Base(bjmPath))
	if err != nil {
		mylog.LogDebug("文件上传失败", err)
	}

	// 3.阿里云录音文件识别
	AudioResult := a.AppConfig.AliYunCloud.AliYunAudioRecognition(remoteUrl, a.AppConfig.IntelligentBlock)

	// 4.生成字幕文件
	a.AppConfig.AliYunCloud.AliyunAudioResultMakeSubtitleFile(videoPath, AudioResult)

}
