package app

import (
	"fmt"
	"io/ioutil"
	"myTool/common"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (a *App)getVideos() []string  {
	files, err := file.GetCurrentFiles(a.AppDir + "/video")
	if err != nil {
		return nil
	}

	var videos []string
	for _, f := range files {
		if ffmpeg.IsVideo(f) {
			videos = append(videos,f)
		}
	}

	return videos

}


func (a *App)ClearTemp() {
	file.RemoveAllFiles(a.TempDir)

	files,_ := file.GetAllFiles(a.AppDir)
	for _, f := range files {
		if strings.HasSuffix(f,".DS_Store") {
			os.Remove(f)
		}
	}
}


func StringToValue(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return v
	}

	if strings.HasPrefix(str, "+") {
		str = strings.TrimPrefix(str, "+")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}

	if strings.HasPrefix(str, "-") {
		str = strings.TrimPrefix(str, "-")
		v, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return v
	}
	return 0
}

// 获取字幕文件
func getSrtPath(f string) string  {
	outputDir := filepath.Dir(f)
	fileSuf := file.GetFileBaseName(f)
	path1 := outputDir + "/" + fileSuf + "_"  + "0.srt"

	if file.PathExist(path1) {
		return path1
	}

	path1 = outputDir + "/" + fileSuf + "_"  + "1.srt"

	if file.PathExist(path1) {
		return path1
	}
	path1 = outputDir + "/" + fileSuf  + ".txt"
	if file.PathExist(path1) {
		return path1
	}

	fmt.Println("字幕文件或者文本文件不存在！")

	return ""

}

func getVoicePath(f string) string  {
	outputDir := filepath.Dir(f)
	fileSuf := file.GetFileBaseName(f)
	path := outputDir + "/" + fileSuf  + ".mp3"
	return path
}

// 获取字幕文件内容
func getSrtContent(f string) string  {

	//字幕文件
	if strings.HasSuffix(f, "srt") {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			return ""
		}

		arr := strings.Split(string(buf), "\n")
		var contents []string

		for _, s := range arr {
			if common.IsChinese(s) {
				contents = append(contents, s)
			}
		}

		return strings.Join(contents,",")

	} else if strings.HasSuffix(f, "txt") {

	}

	return "nil"

}