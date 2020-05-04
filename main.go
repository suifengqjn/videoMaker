package main

import (
	"flag"
	"myProject/videoMaker/GUI"
	"myProject/videoMaker/account"
	"myProject/videoMaker/common"
)

var (
	addr     = flag.String("addr", "127.0.0.1:3435", "address to start the server on")
	appName  = flag.String("appName", "main", "Gowut app name")
	autoOpen = flag.Bool("autoOpen", true, "auto-open the demo in default browser")
)
func main()  {

	appId := account.LoadAppId()
	if len(appId) > 0 {
		account.NewAccount(appId,"")
	}

	cli := common.NewCliEngine()
	eng := common.NewMakerEngine(cli, account.AppAccount)


	//清除aliyun oss 文件
	go eng.ClearRemoteCache()
	go eng.ClearTemp()


	flag.Parse()
	GUI.StartServer(*appName, *addr, *autoOpen)

}
