package account

import (
	"io/ioutil"
	"myTool/appAccount"
	"myTool/common"
	"myTool/file"
	"myTool/sys"
)

var AppAccount = &Account{}

type Account struct {
	*appAccount.AppAccount
}

const (
	AccTypeBase = 1
	AccTypeMonth = 2
	AccTypeYear = 3
)



func NewAccount(appId, inviteCode string) *Account  {
	deviceId := sys.GetSysInfo().DeviceId
	if len(deviceId) > 10 {
		deviceId = deviceId[:10]
	}
	acc := appAccount.NewAppAccount("video_maker", Version,appId, inviteCode, deviceId,false)
	AppAccount = &Account{acc}
	return AppAccount
}

func (e *Account)IsActive() bool  {
	if e == nil || e.AppAccount == nil {
		return false
	}
	return e.AppAccount.AccType > 0
}

func (e *Account)TYPE()string  {
	if e.AppAccount == nil {
		return ""
	}
	if e.AppAccount.AccType == AccTypeBase {
		return "体验版"
	} else if e.AppAccount.AccType == AccTypeMonth {
		return "月卡套餐"
	}  else if e.AppAccount.AccType == AccTypeYear {
		return "年卡套餐"
	}  else {
		return "未知"
	}
}

func (e *Account)Message()string  {
	if e.AppAccount == nil {
		return ""
	}
	return e.AppAccount.Msg
}

func (e *Account)Time()string  {
	if e.AppAccount == nil {
		return ""
	}
	return e.AppAccount.Time
}

func (e *Account)Tip()string  {
	if e.AppAccount == nil {
		return ""
	}
	return e.AppAccount.Tip
}

func (e *Account)InviteCode()(int, string)  {
	if e.AppAccount == nil {
		return 0, ""
	}
	return e.AppAccount.InviteNum, e.AppAccount.InViteCode
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

