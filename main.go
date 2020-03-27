package main

import (
	"flag"
	"myProject/videoMaker/GUI"
	"myProject/videoMaker/account"
	"myProject/videoMaker/app"
)

var (
	addr     = flag.String("addr", "127.0.0.1:3435", "address to start the server on")
	appName  = flag.String("appName", "main", "Gowut app name")
	autoOpen = flag.Bool("autoOpen", true, "auto-open the demo in default browser")
)
func main()  {

	App := app.NewApp()
	App.ClearTemp()
	appId := account.LoadAppId()
	if len(appId) > 0 {
		acc := account.NewAccount(appId,"")
		App.Account = acc
	}

	//清除aliyun oss 文件
	go App.ClearRemoteCache()

	flag.Parse()
	GUI.StartServer(*appName, *addr, *autoOpen)

}
