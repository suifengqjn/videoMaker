package GUI

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"myProject/videoMaker/common"
	"strings"
)

var (
	EndpointTb gwu.TextBox
	AccessKeyIdTb gwu.TextBox
	AccessKeySecretTb gwu.TextBox
	BucketNameTb gwu.TextBox
	BucketDomainTb gwu.TextBox
	ExpirationTb gwu.TextBox

)


func buildPlatform(event gwu.Event) gwu.Comp {
	p := gwu.NewVerticalPanel()

	buildAliYun(p)
	buildAliYunVoice(p)

	buildSaveBtn(p)
	return p
}

func buildAliYun(p gwu.Panel)  {

	row := gwu.NewHorizontalPanel()

	title := gwu.NewLabel("阿里云参数")
	title.Style().SetColor(gwu.ClrBlue)
	row.Add(title)

	row.AddHSpace(20)
	link := gwu.NewLink("参数获取地址", "https://www.aliyun.com/")
	row.Add(link)

	row.AddHSpace(20)
	link = gwu.NewLink("参数教程", "https://github.com/suifengqjn/videoMaker/blob/master/shipin.md")
	row.Add(link)

	p.Add(row)
	buildAliYunOss(p)
}

func buildAliYunOss(p gwu.Panel)  {

	p.AddVSpace(20)
	content := gwu.NewLabel("对象存储OSS")
	content.Style().SetFontSize("20")
	p.Add(content)

	p.AddVSpace(10)
	row1 := gwu.NewHorizontalPanel()
	row1.Add(gwu.NewLabel("AccessKeyId:"))
	row1.AddHSpace(39)
	tb1 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunOss.AccessKeyId)
	tb1.Style().SetWidth("300")
	tb1.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.AccessKeyId = strings.TrimSpace(tb1.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(tb1)
	p.Add(row1)

	p.AddVSpace(5)
	row2 := gwu.NewHorizontalPanel()
	row2.Add(gwu.NewLabel("AccessKeySecret:"))
	row2.Style().SetWidthPx(200)
	row2.AddHSpace(5)
	tb2 := gwu.NewPasswBox(common.MakerEngine.MakerCli.AliYunOss.AccessKeySecret)
	tb2.Style().SetWidth("300")
	tb2.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.AccessKeySecret = strings.TrimSpace(tb2.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(tb2)
	p.Add(row2)

	p.AddVSpace(5)
	row3 := gwu.NewHorizontalPanel()
	row3.Add(gwu.NewLabel("Endpoint:"))
	row3.AddHSpace(66)
	tb3 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunOss.Endpoint)
	tb3.Style().SetWidth("300")
	tb3.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.Endpoint = strings.TrimSpace(tb3.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row3.Add(tb3)
	p.Add(row3)

	p.AddVSpace(5)
	row4 := gwu.NewHorizontalPanel()
	row4.Add(gwu.NewLabel("BucketName:"))
	row4.AddHSpace(36)
	tb4 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunOss.BucketName)
	tb4.Style().SetWidth("300")
	tb4.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.BucketName = strings.TrimSpace(tb4.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row4.Add(tb4)
	p.Add(row4)


	p.AddVSpace(5)
	row5 := gwu.NewHorizontalPanel()
	row5.Add(gwu.NewLabel("BucketDomain:"))
	row5.AddHSpace(22)
	tb5 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunOss.BucketDomain)
	tb5.Style().SetWidth("300")
	tb5.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.BucketDomain = strings.TrimSpace(tb5.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row5.Add(tb5)
	p.Add(row5)

	p.AddVSpace(5)
	row6 := gwu.NewHorizontalPanel()
	row6.Add(gwu.NewLabel("过期时间:"))
	row6.AddHSpace(62)
	tb6 := gwu.NewTextBox(fmt.Sprintf("%v",common.MakerEngine.MakerCli.AliYunOss.Expiration))
	tb6.Style().SetWidth("300")
	tb6.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunOss.Expiration = common.IntValue(tb6.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row6.Add(tb6)
	row6.AddHSpace(10)
	desc := gwu.NewLabel("如比设置3，删除3天前上传到OSS的文件，也可以不设置")
	desc.Style().SetColor(gwu.ClrGray)
	row6.Add(desc)
	p.Add(row6)

}

func buildAliYunVoice(p gwu.Panel) {

	p.AddVSpace(20)
	content := gwu.NewLabel("智能语音交互")
	content.Style().SetFontSize("20")
	p.Add(content)

	p.AddVSpace(10)
	row1 := gwu.NewHorizontalPanel()
	row1.Add(gwu.NewLabel("AccessKeyId:"))
	row1.AddHSpace(39)
	tb1 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunCloud.AccessKeyId)
	tb1.Style().SetWidth("300")
	tb1.AddEHandlerFunc(func(e gwu.Event) {

		common.MakerEngine.MakerCli.AliYunCloud.AccessKeyId = strings.TrimSpace(tb1.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row1.Add(tb1)
	p.Add(row1)

	p.AddVSpace(5)
	row2 := gwu.NewHorizontalPanel()
	row2.Add(gwu.NewLabel("AccessKeySecret:"))
	row2.AddHSpace(5)
	tb2 := gwu.NewPasswBox(common.MakerEngine.MakerCli.AliYunCloud.AccessKeySecret)
	tb2.Style().SetWidth("300")
	tb2.AddEHandlerFunc(func(e gwu.Event) {

		common.MakerEngine.MakerCli.AliYunCloud.AccessKeySecret = strings.TrimSpace(tb2.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(tb2)
	p.Add(row2)

	p.AddVSpace(5)
	row3 := gwu.NewHorizontalPanel()
	row3.Add(gwu.NewLabel("AppKey:"))
	row3.AddHSpace(71)
	tb3 := gwu.NewTextBox(common.MakerEngine.MakerCli.AliYunCloud.AppKey)
	tb3.Style().SetWidth("300")
	tb3.AddEHandlerFunc(func(e gwu.Event) {
		common.MakerEngine.MakerCli.AliYunCloud.AppKey = strings.TrimSpace(tb3.Text())
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row3.Add(tb3)
	p.Add(row3)


}



func buildBaiduFanYi(p gwu.Panel) {
	//
	//srt := app.Engine.GetSrtConf()
	//
	//p.AddVSpace(20)
	//content := gwu.NewLabel("百度翻译")
	//content.Style().SetFontSize("20")
	//p.Add(content)
	//
	//p.AddVSpace(10)
	//row1 := gwu.NewHorizontalPanel()
	//row1.Add(gwu.NewLabel("APP Id:"))
	//row1.AddHSpace(39)
	//tb1 := gwu.NewTextBox(srt.AliYunCloud.AccessKeyId)
	//tb1.Style().SetWidth("300")
	//tb1.AddEHandlerFunc(func(e gwu.Event) {
	//
	//	srt.AliYunCloud.AccessKeyId = strings.TrimSpace(tb1.Text())
	//}, gwu.ETypeChange, gwu.ETypeKeyUp)
	//row1.Add(tb1)
	//p.Add(row1)
	//
	//p.AddVSpace(5)
	//row2 := gwu.NewHorizontalPanel()
	//row2.Add(gwu.NewLabel("APP Secret:"))
	//row2.AddHSpace(5)
	//tb2 := gwu.NewPasswBox(srt.AliYunCloud.AccessKeySecret)
	//tb2.Style().SetWidth("300")
	//tb2.AddEHandlerFunc(func(e gwu.Event) {
	//
	//	srt.AliYunCloud.AccessKeySecret = strings.TrimSpace(tb2.Text())
	//}, gwu.ETypeChange, gwu.ETypeKeyUp)
	//row2.Add(tb2)
	//p.Add(row2)


}

var checkLabel gwu.Label
func buildSaveBtn(p gwu.Panel)  {

	p.AddVSpace(20)

	line := gwu.NewHorizontalPanel()

	saveBtn := gwu.NewButton("    保存   ")
	saveBtn.Style().SetFontSize("18")
	saveBtn.Style().SetColor(gwu.ClrBlue)

	saveBtn.AddEHandlerFunc(func(e gwu.Event) {
		common.SavePlatFormParam()

	}, gwu.ETypeClick)

	line.Add(saveBtn)

	line.AddHSpace(20)
	checkBtn := gwu.NewButton("    参数检测   ")
	checkBtn.Style().SetFontSize("18")
	checkBtn.Style().SetColor(gwu.ClrBlue)

	checkBtn.AddEHandlerFunc(func(e gwu.Event) {
		path := "./source/files/test.txt"
		err := common.CheckAliYun(path)
		if err == nil {
			checkLabel.SetText("参数正确")
		} else {
			checkLabel.SetText("参数错误, 请查看控制台输出")
		}
		e.MarkDirty(checkLabel)

	}, gwu.ETypeClick)

	line.Add(checkBtn)

	line.AddHSpace(20)
	checkLabel = gwu.NewLabel("")
	checkLabel.Style().SetFontSize("20")
	line.Add(checkLabel)
	p.Add(line)
}