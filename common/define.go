package common

var (
	appIdPath   = "./source/files/app_id.txt"
	configPath  = "./source/files/conf.json"
	config2Path = "./source/files/config2.json"
	config3Path = "./source/files/config3.json"

	imagePath = "./material"
	filmPath  = "./material"
	fontPath  = "./source/font"
	confPath  = "./source/files/conf.json"

	SpitCount = 200 //长文本分段个数
)

func AppKeyPath() string {
	return appIdPath
}
