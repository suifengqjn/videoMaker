package main

import (
	"flag"
	"fmt"
	"myProject/videoMaker/GUI"
	"myProject/videoMaker/app"
)
var (
	addr     = flag.String("addr", "127.0.0.1:3435", "address to start the server on")
	appName  = flag.String("appName", "main", "Gowut app name")
	autoOpen = flag.Bool("autoOpen", true, "auto-open the demo in default browser")
)
func main()  {

	App := app.NewApp()

	fmt.Println(App)


	path := "/Users/qianjianeng/Desktop/annie/test_video1.mp4"

	//App.GenerateSrt(path)
	fmt.Println(App, path)

	flag.Parse()
	GUI.StartServer(*appName, *addr, *autoOpen)

}
