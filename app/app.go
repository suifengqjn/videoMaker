package app

import (
	"myProject/videoMaker/account"
	"myProject/videoMaker/common"
)

type App struct {
	AppConfig *common.Config
	Account *account.Account
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