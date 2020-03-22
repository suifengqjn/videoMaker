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
	bjmPath := ffmpeg.ExtractBgm(a.FCmd, videoPath)

	// 2. 上传到oss
	remoteUrl, err := a.AliYunOss.UploadObject(bjmPath, filepath.Base(bjmPath))
	if err != nil {
		mylog.LogDebug("文件上传失败", err)
	}

	// 3.阿里云录音文件识别
	AudioResult := a.AliYunCloud.AliYunAudioRecognition(remoteUrl, a.IntelligentBlock)

	// 4.生成字幕文件
	a.AliYunCloud.AliyunAudioResultMakeSubtitleFile(videoPath, AudioResult)

}

func (a *App)createSrtWithUrl(videoPath,url string)  {
	// 3.阿里云录音文件识别
	AudioResult := a.AliYunCloud.AliYunAudioRecognition(url, a.IntelligentBlock)

	// 4.生成字幕文件
	a.AliYunCloud.AliyunAudioResultMakeSubtitleFile(videoPath, AudioResult)
}
