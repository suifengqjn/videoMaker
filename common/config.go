package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myTool/aliyun/cloud"
	"myTool/aliyun/oss"
	"myTool/common"
	"myTool/file"
	"myTool/sys"
	"os"
)

var AppConfig *Config

type Config struct {
	*SrtConfig
	*GUIConfig `json:"-"`
}

type SrtConfig struct {
	Setting setting
	FCmd string
	AliYunOss *oss.AliYunOss //oss
	AliYunCloud *cloud.AliYunCloud  //语音识别引擎
	IntelligentBlock bool //智能分段处理
	TempDir string //临时文件目录
	AppDir string //应用根目录
	OutPutDir string //输出目录
}

type GUIConfig struct {
	CutFront CutFront
	CutBack CutBack
	ClearWater ClearWater
	ClearWater1 ClearWater
	ExtractSubtitles ExtractSubtitles
	Composite Composite
	Subtitles Subtitles
	WaterText WaterText
	RunWaterText RunWaterText
	WaterImage WaterImage
	AddBgm AddBgm
	FilmHead FilmTitle
	FilmFoot FilmEnd
}

type setting struct {
	CurrentEngineId int //目前使用引擎Id
	MaxConcurrency int //任务最大处理并发数
	OutputType int //输出文件类型
	OutputEncode int //输出文件编码
	SrtFileDir string //Srt文件输出目录
	SoundTrack int //输出音轨
}

func NewAppConfig() *Config  {
	projectDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法识别路径")
		panic(err)
	}

	var oss = &oss.AliYunOss{}
	var clo =&cloud.AliYunCloud{}
	conf, err := LoadSrtConf()
	if err == nil {
		oss = conf.AliYunOss
		clo = conf.AliYunCloud
	}

	srt := &SrtConfig{
		Setting:          setting{},
		FCmd:             GetFCmd(0),
		AliYunOss:        oss,
		AliYunCloud:      clo,
		IntelligentBlock: true,
		TempDir:          projectDir + "/source/tempVideos",
		AppDir:           projectDir,
		OutPutDir:        file.GetDeskTop() + "/video_maker",
	}
	AppConfig = &Config{
		srt,
		&GUIConfig{},
	}
	return AppConfig

}



func GetFCmd(system int) string {

	if system == 0 {
		info := sys.GetSysInfo()
		system = info.PlatForm
	}
	if system == sys.MacOS {
		return "./source/mac/tool"
	} else if system == sys.Win64 {
		return "./source/win/64/tool.exe"
	} else if system == sys.Win32 {
		return "./source/win/32/tool.exe"
	}
	return ""

}

func LoadSrtConf()(*Config, error)  {

	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(buf,&conf)
	if err != nil {
		return nil,err
	}
	return &conf, nil
}

func SaveSrtConf()  {
	buf, err := json.Marshal(AppConfig)
	if err != nil {
		return
	}
	common.CoverWriteToFile(confPath, buf)

}
