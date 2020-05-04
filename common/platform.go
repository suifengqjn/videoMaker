package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	cm "myProject/videoCli/common"
	"myTool/aliyun/cloud"
	"myTool/file"
)

func SavePlatFormParam() {

	buf, err := json.Marshal(MakerEngine.MakerCli.PlatformConf)
	if err != nil {
		file.OverWrite(PlatFormParamPath, string(buf))
	}

}

func LoadPlatFormParam() *cm.PlatformConf {
	buf, err := ioutil.ReadFile(PlatFormParamPath)
	if err != nil {
		return nil
	}
	var conf cm.PlatformConf
	err = json.Unmarshal(buf, &conf)
	if err != nil {
		return nil
	}

	return &conf
}

func CheckAliYun(textPath string) error {

	// 获取文本内容
	buf, err := ioutil.ReadFile(textPath)
	if err != nil {
		return err
	}
	content := string(buf)
	if len(content) == 0 {
		fmt.Println("文件内容为空, 检查此文件", textPath)
		return errors.New("检测文件内容为空")
	}

	fmt.Println("进行参数测试...")
	// 语音输出路径
	//output := dir + "/" + cm.GetRandomString(6) + ".mp3"
	param := cloud.TTSParam{
		Voice:      "",
		Volume:     50,
		SpeechRate: 0,
		PitchRate:  0,
	}

	res, err := MakerEngine.MakerCli.TTSVoiceMerge(content,MakerEngine.MakerCli.TempDir, &param)
	if err != nil || file.PathExist(res) == false {
		fmt.Println("测试失败！", err)
		return err
	}
	fmt.Println("参数填写正确！！！")
	return nil
}
