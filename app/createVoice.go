package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"myTool/aliyun/cloud"
	cm "myTool/common"
	"net/http"
)

func (a *App) createVoice(dir string) (string, error) {
	// 获取文本文件
	textPath := getTextOrSrtPath(dir)

	// 获取文本内容
	content := getSrtContent(textPath)

	if len(content) == 0{
		fmt.Println("文件内容为空, 检查此目录", dir)
		return "", errors.New("err")
	}

	fmt.Println("进行文本转语音...")
	// 语音输出路径
	output := dir + "/" + cm.GetRandomString(6) + ".mp3"
	param := cloud.TTSParam{
		Voice:      a.Composite.Voice,
		Volume:     a.Composite.Volume,
		SpeechRate: a.Composite.SpeechRate,
		PitchRate:  a.Composite.PitchRate,
	}

	path, err := a.TTSVoice(content, output, &param)
	if err != nil  {
		fmt.Println("文本转语音失败！",err)
		return "", err
	}
	return path, nil
}

func (a *App)TTSVoice(content, output string,param *cloud.TTSParam)(string, error)  {
	path, err := a.AliYunCloud.TTSToVoice(content, output, param)
	if err != nil || len(path) == 0 {
		fmt.Println("文本转语音失败")
	}

	return path, err
}

func (a *App)LongTTSVoice(content, output string,param *cloud.TTSParam)(string, error)  {
	//长文本
	url, err := a.AliYunCloud.LongTTSToVoice(content, output, param)
	if err != nil || len(url) == 0 {
		fmt.Println("文本转语音失败")
	}

	res, err := http.Get(url)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	cm.CoverWriteToFile(output, buf)
	return url, nil
}

