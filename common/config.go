package common

import (
	"fmt"
	cm "myProject/videoCli/common"
	"myProject/videoCli/makerCli"
	"myTool/file"
	"myTool/fmg"
	"myTool/sys"
	"os"
)

var AppConfig *cm.MakerConfig
var cliEngine *makerCli.Engine

func NewCliEngine() *makerCli.Engine {

	callback := func(msg string, video string) {
		fmt.Println(video, msg)
	}
	fg := fmg.NewFmg(GetFCmd(0),
		"./source/simsun.ttc",
		"./source/tempVideos",
		"source",
		callback,
	)

	cliEngine = makerCli.NewMakerEngineCli(fg, NewAppConfig())
	return cliEngine
}

func NewAppConfig() *cm.MakerConfig {
	conf := LoadPlatFormParam()
	if conf == nil {
		conf = cm.NewPlatformConf()
	}

	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	outPut := file.GetDeskTop() + "/videoMakerOut"
	if file.PathExist(outPut) == false {
		os.MkdirAll(outPut, os.ModePerm)
	}
	AppConfig = &cm.MakerConfig{
		PlatformConf: conf,
		AppConf:      cm.AppConf{},
		Setting: cm.Setting{
			ProjectDir: cur,
			WorkDir:    "./video",
			OutPutDir:  outPut,
		},
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

//
//func LoadAppConf() (*cm.MakerConfig, error) {
//
//	buf, err := ioutil.ReadFile(confPath)
//	if err != nil {
//		return nil, err
//	}
//	var conf cm.MakerConfig
//	err = json.Unmarshal(buf, &conf)
//	if err != nil {
//		return nil, err
//	}
//	return &conf, nil
//}
//
//func SaveAppConf() {
//	buf, err := json.Marshal(AppConfig)
//	if err != nil {
//		return
//	}
//	common.CoverWriteToFile(confPath, buf)
//
//}
