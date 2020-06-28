package account

import (
	"io/ioutil"
	"myTool/appAccount"
	"myTool/common"
	"myTool/file"
)

var AppAccount  *appAccount.AppAccount

func NewAccount(appId, inviteCode string) *appAccount.AppAccount  {

	AppAccount = appAccount.NewAccount(dbName,appId, inviteCode, Version)
	return AppAccount
}

func LoadAppId() string  {
	if file.PathExist(AppKeyPath()) {
		byte, err := ioutil.ReadFile(AppKeyPath())
		if err != nil || len(byte) < 10 {
			return ""
		}
		return string(byte)
	}
	return ""
}

func SaveAppId(appId string)  {
	common.CoverWriteToFile(AppKeyPath(), []byte(appId))
}


func AppKeyPath() string {
	return "./source/files/app_id.txt"
}

