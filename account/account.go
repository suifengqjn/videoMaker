package account

import "myTool/appAccount"

var Acc *appAccount.AppAccount

func NewAccount() *appAccount.AppAccount  {

	Acc = appAccount.NewAppAccount("video_maker", "1.0","", false)
	return Acc

}

