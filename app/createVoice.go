package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"myProject/videoMaker/common"
	"myTool/aliyun/cloud"
	cm "myTool/common"
	"myTool/ffmpeg"
	"myTool/file"
	"net/http"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func (a *App) createVoice(dir string) (string, error) {
	// 获取文本文件
	textPath := getTextPath(dir)

	// 获取文本内容
	content := getSrtContent(textPath)

	if len(content) == 0 {
		fmt.Println("文件内容为空, 检查此目录", dir)
		return "", errors.New("err")
	}

	fmt.Println("进行文本转语音...")
	// 语音输出路径
	//output := dir + "/" + cm.GetRandomString(6) + ".mp3"
	param := cloud.TTSParam{
		Voice:      a.Composite.Voice,
		Volume:     a.Composite.Volume,
		SpeechRate: a.Composite.SpeechRate,
		PitchRate:  a.Composite.PitchRate,
	}

	path, err := a.TTSVoiceMerge(content, dir, &param)
	if err != nil  {
		fmt.Println("文本转语音失败！",err)
		return "", err
	}
	return path, nil
}

//当字数超过 ~300 ，分段之后合成
func (a *App)TTSVoiceMerge(content, dir string,param *cloud.TTSParam) (string, error) {

	arr := a.SpitCount(content)
	var spanVoices []string
	var out string
	if len(arr) == 1 {
		output := dir + "/" + cm.GetRandomString(6) + ".mp3"
		return a.TTSVoice(content, output, param)
	} else {

		for _, s := range arr {
			output := ffmpeg.MakeRandExportPath("mp3")
			path, err := a.TTSVoice(s, output, param)
			if err != nil {
				return "", err
			}
			spanVoices = append(spanVoices, path)
		}

		//合并音频片段
		output := dir + "/" + cm.GetRandomString(6) + ".mp3"
		output, err := ffmpeg.MergeBgms(a.FCmd, spanVoices,output)
		if err != nil {
			return "", err
		}
		out = output

	}

	return out,nil
}

// 文字分段
func (a *App)SpitCount(content string) []string  {
	var arr []string
	if cm.IsChinese(content) {
		count := utf8.RuneCountInString(content)
		if count > common.SpitCount {

			sp := strings.Split(content,",")
			lines := len(sp)
			temp := ""
			var tempArr []string
			for i := 0;i < lines;i ++ {
				temp += sp[i]
				tempArr = append(tempArr, sp[i])
				if utf8.RuneCountInString(temp) >= common.SpitCount {
					arr = append(arr,strings.Join(tempArr, ","))
					temp = ""
					tempArr = tempArr[0:0]
				}
				if i == lines - 1  && utf8.RuneCountInString(temp) < common.SpitCount {
					if len(temp) > 0 {
						arr = append(arr,strings.Join(tempArr, ","))
					}
				}

			}

		} else {
			arr = append(arr,content)
		}
	} else {  //英文

		if len(content) > 260 {

			sp := strings.Split(content,",")
			lines := len(sp)
			temp := ""
			var tempArr []string
			for i := 0;i < lines;i ++ {
				temp += sp[i]
				tempArr = append(tempArr, sp[i])
				if len(temp) >= 260 {
					arr = append(arr,strings.Join(tempArr, ","))
					temp = ""
					tempArr = tempArr[0:0]
				}
				if i == lines - 1  && utf8.RuneCountInString(temp) < common.SpitCount {
					if len(temp) > 0 {
						arr = append(arr,strings.Join(tempArr, ","))
					}
				}

			}

		} else {
			arr = append(arr,content)
		}

	}

	//tempArr := arr
	////如果最后两个之和小于220，则合并成一个
	//if len(arr) >= 2 {
	//	if utf8.RuneCountInString(arr[len(arr)-1]) +  utf8.RuneCountInString(arr[len(arr)-2]) < 220{
	//		temp := arr[len(arr)-2] + "," + arr[len(arr)-1]
	//		arr = arr[:len(arr)-2]
	//		arr = append(arr, temp)
	//	}
	//}
	//fmt.Println(tempArr)

	return arr
}

// 中文 261 个字
// 英文300 个字符
// //"<speak>请闭上眼睛休息一下<break time=\"500ms\"/>好了，请睁开眼睛。<break time=\"600ms\"/></speak>"
func (a *App)TTSVoice(content, output string,param *cloud.TTSParam)(string, error)  {

	newContent := a.formatSSML(content)
	path, err := a.AliYunCloud.TTSToVoicePOST(newContent, output, param)
	if err != nil || len(path) == 0 {
		fmt.Println("文本转语音失败")
	}

	return path, err
}

// 只有中文支持
func (a *App)formatSSML(content string) string  {
	arr := strings.Split(content,`,`)
	newContent := `<speak>`
	for _, s := range arr {
		newContent = newContent + s + fmt.Sprintf(`<break time="%vms"/>`, a.Composite.BreakTime)
	}
	newContent = newContent + `</speak>`
	return newContent
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

func (a *App)CheckAliYun(textPath string) error  {

	// 获取文本内容
	content := getSrtContent(textPath)

	if len(content) == 0 {
		fmt.Println("文件内容为空, 检查此文件", textPath)
		return  errors.New("检测文件内容为空")
	}

	fmt.Println("进行参数测试...")
	// 语音输出路径
	//output := dir + "/" + cm.GetRandomString(6) + ".mp3"
	param := cloud.TTSParam{
		Voice:      a.Composite.Voice,
		Volume:     a.Composite.Volume,
		SpeechRate: a.Composite.SpeechRate,
		PitchRate:  a.Composite.PitchRate,
	}

	res, err := a.TTSVoiceMerge(content, filepath.Dir(textPath), &param)
	if err != nil || file.PathExist(res) == false {
		fmt.Println("文本转语音失败！",err)
		return err
	}
	fmt.Println("参数填写正确！！！")
	return nil
}