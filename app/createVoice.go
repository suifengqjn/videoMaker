package app

import (
	"fmt"
	"io/ioutil"
	"myTool/aliyun/cloud"
	cm "myTool/common"
	"net/http"
)

func (a *App) createVoice(videoPath string) (url ,path string) {
	// 获取文件
	textPath := getSrtPath(videoPath)
	fmt.Println(textPath)

	// 获取文本内容
	content := getSrtContent(textPath)
	fmt.Println(content)

	// 文本转语音
	output := getVoicePath(videoPath)
	param := cloud.TTSParam{
		Voice:      a.TextToVoice.Voice,
		Volume:     a.TextToVoice.Volume,
		SpeechRate: a.TextToVoice.SpeechRate,
		PitchRate:  a.TextToVoice.PitchRate,
	}
	url, err := a.AliYunCloud.LongTTSToVoice(content, output, &param)
	if err != nil || len(url) == 0 {
		fmt.Println("文本转语音失败")
	}

	res, err := http.Get(url)
	if err != nil {
		return "",""
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	cm.CoverWriteToFile(output, buf)

	return url, output
}
