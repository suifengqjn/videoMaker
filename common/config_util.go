package common

import (
	"encoding/json"
	"io/ioutil"
	cm "myProject/videoCli/common"
	"myProject/videoMaker/common"
)
var (
	config1Path = "./source/files/config1.json"
	config2Path = "./source/files/config2.json"
	config3Path = "./source/files/config3.json"
)


func LoadConfig1() *cm.MakerConfig  {
	buf, err := ioutil.ReadFile(config1Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &common.AppConfig)
	if err != nil {
		return nil
	}

	return common.AppConfig

}


func LoadConfig2() *cm.MakerConfig  {
	buf, err := ioutil.ReadFile(config2Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &common.AppConfig)
	if err != nil {
		return nil
	}

	return common.AppConfig

}

func LoadConfig3() *cm.MakerConfig  {
	buf, err := ioutil.ReadFile(config3Path)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(buf, &common.AppConfig)
	if err != nil {
		return nil
	}

	return common.AppConfig

}
func SaveConfig1(con *cm.MakerConfig) {

	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config1Path, buf)
	}

}

func SaveConfig2(con *cm.MakerConfig) {
	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config2Path, buf)
	}
}

func SaveConfig3(con *cm.MakerConfig) {
	buf, err := json.Marshal(con)
	if err == nil {
		writeToFile(config3Path, buf)
	}
}



