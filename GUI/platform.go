package GUI

import "github.com/icza/gowut/gwu"


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

	// oss


	//语音参数



	return p
}

/*
	Endpoint string
	AccessKeyId string
	AccessKeySecret string
	BucketName string //yourBucketName
	BucketDomain string //Bucket 域名
	Expiration int // 设置几天前的文件删除
*/
func buildOssUI(p gwu.Panel) {
	row := gwu.NewVerticalPanel()

	row1 := gwu.NewHorizontalPanel()
	content := gwu.NewLabel("Endpoint")
	row1.Add(content)

	row.Add(row1)



	p.Add(row)
}