package app

import (
	"myProject/videoMaker/account"
	"myProject/videoMaker/common"
	"myTool/appAccount"
)

type App struct {
	AppConfig *common.Config
	Account *appAccount.AppAccount
}

func NewApp() *App  {
	app := &App{
		AppConfig: common.NewAppConfig(),
		Account:   account.NewAccount(),
	}
	return app
}

func (a *App)GenerateSrt(videoPath string)  {
	a.createSrt(videoPath)
}