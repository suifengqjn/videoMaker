package common

import (
	"myTool/file"
	"os"
	"path/filepath"
)

var (
	configPath  = "./source/files/conf.json"
	config2Path = "./source/files/config2.json"
	config3Path = "./source/files/config3.json"

	imagePath = "./material"
	filmPath  = "./material"
	fontPath  = "./source/font"
	confPath  = "./source/files/conf.json"
)

func AppKeyPath() string {
	dir := file.GetHomeDir() + "/.video_app_key"
	if file.PathExist(dir) == false {
		os.MkdirAll(dir, os.ModePerm)
	}
	return filepath.Join(dir, "video_maker_app_id.txt")
}
