package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)


var (
	config1Path = "./source/files/config1.json"
	config2Path = "./source/files/config2.json"
	config3Path = "./source/files/config3.json"

	imagePath = "./material"
	filmPath  = "./material"
	fontPath  = "./source/font"
)

var (
	imageSufs = []string{"jpg", "jpeg", "png"}
	fontSufs  = []string{"ttc", "ttf"}
)

var (
	TextColors    = []string{"未选择", "red", "yellow", "green", "black", "white", "blue", "gray", "purple"}
	TextColorsMap = map[string]int{
		"red":    1,
		"yellow": 2,
		"green":  3,
		"black":  4,
		"white":  5,
		"blue":   6,
		"gray":   7,
		"purple": 8,
	}

	WaterStyle    = []string{"未选择", "左上角", "右上角", "右下角", "左下角", "正中间", "顶部居中", "右侧居中", "底部居中", "左侧居中"}
	WaterStyleMap = map[int]string{
		1: "左上角",
		2: "右上角",
		3: "右下角",
		4: "左下角",
		5: "正中间",
		6: "顶部居中",
		7: "右侧居中",
		8: "底部居中",
		9: "左侧居中"}

	PinPStyle    = []string{"未选择", "左上角", "右上角", "右下角", "左下角"}
	PinPStyleMap = map[int]string{
		1: "左上角",
		2: "右上角",
		3: "右下角",
		4: "左下角"}

	BjImageStyle    = []string{"未选择", "自动", "左右", "上下"}
	BjImageStyleMap = map[int]string{
		1: "自动",
		2: "左右",
		3: "上下",
	}
)

func LoadVideos() []string {
	var videos []string
	files, dirs, err := file.GetCurrentFilesAndDirs(AppConfig.AppDir)
	if err == nil {
		videos = append(videos, files...)
	}

	for _, d := range dirs {

		if strings.HasSuffix(d, "video/result") || strings.HasSuffix(d, `video\result`) {
			continue
		}
		// 处理第一级文件夹内的视频
		files, err := file.GetCurrentFiles(d)
		if err != nil {
			continue
		}
		videos = append(videos, files...)

	}

	return videos

}

func StrValue(v interface{}) string {

	if fmt.Sprintf("%v", v) == "0" {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func IntValue(v string) int {
	v = strings.TrimSpace(v)
	value, err := strconv.Atoi(v)
	if err != nil {
		valueF, err := strconv.ParseFloat(v, 10)
		if err != nil {
			return 0
		} else {
			return int(valueF)
		}
	}
	return value
}


func LoadConfig1() *Config {
	buf, err := ioutil.ReadFile(config1Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &AppConfig)
	if err != nil {
		return nil
	}

	return AppConfig

}

func SaveConfig1(con *Config) {

	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config1Path, buf)
	}

}



func writeToFile(filepath string, buf []byte) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err == nil {
		file.Write(buf)
	}
}

func LoadFonts() (map[string]string, []string) {
	files, err := file.GetAllFiles(fontPath)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		baseName := file.GetFileBaseName(f)
		suf := file.GetFileSuf(f)
		if Contains(fontSufs, strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	return res, keys
}

// fileName : filePath
func LoadImages() (map[string]string, []string) {
	files, err := file.GetAllFiles(imagePath)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		baseName := filepath.Base(f)
		suf := strings.Split(baseName, ".")[1]
		if Contains(imageSufs, strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	sort.Strings(keys)
	return res, keys
}

func LoadFilms() (map[string]string, []string) {
	files, err := file.GetAllFiles(filmPath)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]string)
	var keys []string
	for _, f := range files {
		baseName := filepath.Base(f)
		suf := strings.Split(baseName, ".")[1]
		if Contains(ffmpeg.VideoSufs(), strings.ToLower(suf)) {
			res[baseName] = f
			keys = append(keys, baseName)
		}
	}
	sort.Strings(keys)
	return res, keys
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
