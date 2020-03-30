package app

import (
	"fmt"
	"io/ioutil"
	"myTool/common"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"strings"
)

func (a *App)getVideoDirs() []string  {
	_,dirs, err := file.GetCurrentFilesAndDirs(a.AppDir + "/video")
	if err != nil {
		return nil
	}

	var videoDirs []string
	for _, f := range dirs {
		videoDirs = append(videoDirs,f)
	}

	return videoDirs

}

func getVideos(dir string)[]string  {
	files, err := file.GetCurrentFiles(dir)
	if err != nil {
		return nil
	}

	var videoDirs []string
	for _, f := range files {
		if ffmpeg.IsVideo(f) {
			videoDirs = append(videoDirs,f)
		}
	}

	return videoDirs
}

func getImages(dir string)[]string  {
	files, err := file.GetCurrentFiles(dir)
	if err != nil {
		return nil
	}

	var images []string
	for _, f := range files {
		if ffmpeg.IsImage(f) {
			images = append(images,f)
		}
	}

	return images
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


// 获取字幕文件 从文件夹中找出 txt 或 srt
func getTextOrSrtPath(dir string) string  {

	files, _ := file.GetCurrentFiles(dir)
	for _, f := range files {
		if strings.HasSuffix(f,".txt") {
			return f
		}
	}

	for _, f := range files {
		if strings.HasSuffix(f,".srt") {
			return f
		}
	}

	return ""

}

func getSrtPath(dir string) string  {

	files, _ := file.GetCurrentFiles(dir)
	for _, f := range files {
		if strings.HasSuffix(f,".srt") {
			return f
		}
	}

	return ""

}

func getVoicePath(dir string) string  {
	files, _ := file.GetCurrentFiles(dir)
	for _, f := range files {
		if strings.HasSuffix(f,".mp3") {
			return f
		}
	}

	return ""
}

func getVideoPath(dir string) string  {
	files, _ := file.GetCurrentFiles(dir)
	for _, f := range files {
		if ffmpeg.IsVideo(f) {
			return f
		}
	}
	return ""
}

// 获取字幕文件内容

var spitRune = []rune{'，','。','？','！','!','\n',',','.'}
func getSrtContent(f string) string  {

	//字幕文件
	if strings.HasSuffix(f, "srt") {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			return ""
		}

		arr := strings.Split(string(buf), "\n")
		var contents []string
		//"<speak>请闭上眼睛休息一下<break time=\"500ms\"/>好了，请睁开眼睛。</speak>"
		for _, s := range arr {
			if common.IsChinese(s) {
				contents = append(contents, fmt.Sprintf(`%v<break time="500ms"/>`,s))
			}
		}
		return `<speak>` + strings.Join(contents,``) + `</speak>`

	} else if strings.HasSuffix(f, "txt") {

		buf, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Println(f, "文本文件有问题，请检查")
			return ""
		}
		result := strings.FieldsFunc(string(buf), func(c rune) bool {

			for _, r := range spitRune {
				if c == r {
					return true
				}
			}
			return false
		})

		var res []string
		for _, s := range result {
			if len(strings.TrimSpace(s)) > 0 {
				res = append(res, fmt.Sprintf(`%v<break time="500ms"/>`,s))
			}
		}
		return `<speak>` + strings.Join(res,``) + `</speak>`
	}
	return ""
}



func (a *App)BgmDir() string  {
	return a.AppDir + "/bgm"
}

func (a *App)FontPath() string {
	return "./source/simsun.ttc"
}