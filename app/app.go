package app

import (
	"myProject/videoMaker/account"
	"myProject/videoMaker/common"
)


var Engine *App

type App struct {
	*common.Config
	*account.Account
}

func NewApp() *App  {
	Engine = &App{
	common.NewAppConfig(),
	nil,
	}
	return Engine
}



func (a *App)GetSrtConf() *common.SrtConfig  {
	if a == nil {
		return nil
	}
	return a.SrtConfig
}
