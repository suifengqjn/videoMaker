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

	Voiceover = []string{"小云","小刚","若兮","思琪","思佳","思诚","艾琪","艾佳","艾诚","艾达","宁儿","瑞琳","思悦","艾雅","艾夏","艾美","艾雨","艾悦","艾婧","小美","艾娜","伊娜","思婧","思彤","小北","艾彤","艾薇","艾宝","Harry","Abby","Andy","Eric","Emily","Luna","Luca","Wendy","William","Olivia"}

	VoiceoverMap = map[string]string{
		"Abby":"Abby","Andy":"Andy","Emily":"Emily","Eric":"Eric","Harry":"Harry","Luca":"Luca","Luna":"Luna","Olivia":"Olivia","Wendy":"Wendy","William":"William","伊娜":"Yina","宁儿":"Ninger","小云":"Xiaoyun","小刚":"Xiaogang","小北":"Xiaobei","小美":"Xiaomei","思佳":"Sijia","思婧":"Sijing","思彤":"Sitong","思悦":"Siyue","思琪":"Siqi","思诚":"Sicheng","瑞琳":"Ruilin","艾佳":"Aijia","艾夏":"Aixia","艾娜":"Aina","艾婧":"Aijing","艾宝":"Aibao","艾彤":"Aitong","艾悦":"Aiyue","艾琪":"Aiqi","艾美":"Aimei","艾薇":"Aiwei","艾诚":"Aicheng","艾达":"Aida","艾雅":"Aiya","艾雨":"Aiyu","若兮":"Ruoxi",
	}

	CompleteStyle = []string{"未选择","配音加字幕","仅配音","仅字幕"}
	CompleteStyleMap = map[string]int{
		"未选择":0,
		"配音加字幕":1,
		"仅配音":2,
		"仅字幕":3,
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
