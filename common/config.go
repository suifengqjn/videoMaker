package common

import (
	"fmt"
	"myTool/aliyun/cloud"
	"myTool/aliyun/oss"
	"myTool/sys"
	"os"
)

var AppConfig *Config

type Config struct {
	Setting setting
	FCmd string
	AliYunOss *oss.AliYunOss //oss
	AliYunCloud *cloud.AliYunCloud  //语音识别引擎

	CutFront CutFront
	CutBack CutBack
	ClearWater ClearWater
	ClearWater1 ClearWater
	WaterText WaterText
	RunWaterText RunWaterText
	WaterImage WaterImage
	AddBgm AddBgm
	FilmHead FilmTitle
	FilmFoot FilmEnd

	IntelligentBlock bool //智能分段处理
	TempDir string //临时文件目录
	AppDir string //应用根目录
	OutPutDir string //输出目录
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

	oss := &oss.AliYunOss{
		Endpoint:        "oss-cn-beijing.aliyuncs.com",
		AccessKeyId:     "LTAI4Fr1h6k7YcfU7MGERKBB",
		AccessKeySecret: "JkXerf7S2f0IV6TYjS4liLXpUSVo2s",
		BucketName:      "filecloud-store",
		BucketDomain:    "filecloud-store.oss-cn-beijing.aliyuncs.com",
		Expiration:      7,
	}

	clo := &cloud.AliYunCloud{
		AccessKeyId:     "LTAI4Fr1h6k7YcfU7MGERKBB",
		AccessKeySecret: "JkXerf7S2f0IV6TYjS4liLXpUSVo2s",
		AppKey:          "0DaW2ROvgK4ZIpIA",
	}

	AppConfig := &Config{
		FCmd:GetFCmd(0),
		AliYunOss:oss,
		AliYunCloud:clo,
		IntelligentBlock:true,
		TempDir:projectDir + "/source/temp_data",
		AppDir:projectDir,
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