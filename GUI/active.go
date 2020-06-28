package GUI

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"myProject/videoMaker/account"
	"strings"
)
var (
	l1 gwu.Label
	l2 gwu.Label
	l3 gwu.Label
)
func buildActiveUI(event gwu.Event) gwu.Comp {
	
	p := gwu.NewVerticalPanel()

	if account.AppAccount.IsActive() {
		buildActive(p)
	} else {
		buildUnActive(p)
	}

	buildInviteBottom(p)
	return p
	
}

func buildActive(p gwu.Panel)  {
	p.Add(gwu.NewLabel("软件已经激活"))

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("账户信息"))

	p.AddVSpace(20)
	p.Add(gwu.NewLabel(account.AppAccount.TYPE()))
	p.AddVSpace(5)
	p.Add(gwu.NewLabel(account.AppAccount.Message()))
	p.AddVSpace(5)
	p.Add(gwu.NewLabel(account.AppAccount.TimeInfo()))

}

func buildUnActive(p gwu.Panel)  {
	link := gwu.NewLink("卡号购买", "https://pr.kuaifaka.com/item/3ZUpQ")
	link.Style().SetColor("red")
	p.Add(link)

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("填写卡号进行激活, 邀请码没有可以不填"))

	p.AddVSpace(20)

	row := gwu.NewHorizontalPanel()

	row.Add(gwu.NewLabel("卡号:"))
	row.AddHSpace(18)

	tb := gwu.NewTextBox("")
	tb.SetMaxLength(100)
	tb.Style().SetWidthPx(400)
	tb.AddSyncOnETypes(gwu.ETypeKeyUp)
	tb.AddEHandlerFunc(func(e gwu.Event) {

	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(tb)
	p.Add(row)

	p.AddVSpace(10)
	row2 := gwu.NewHorizontalPanel()

	row2.Add(gwu.NewLabel("邀请码:"))
	inviteTb := gwu.NewTextBox("")
	inviteTb.SetMaxLength(50)
	inviteTb.Style().SetWidthPx(400)
	inviteTb.AddSyncOnETypes(gwu.ETypeKeyUp)
	inviteTb.AddEHandlerFunc(func(e gwu.Event) {

	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(inviteTb)
	p.Add(row2)


	p.AddVSpace(20)
	user := gwu.NewLabel("")
	p.Add(user)

	p.AddVSpace(20)
	message := gwu.NewLabel("")
	p.Add(message)

	p.AddVSpace(20)
	timeStr := gwu.NewLabel("")
	p.Add(timeStr)

	p.AddVSpace(20)
	b := gwu.NewButton("     激 活    ")
	b.Style().SetWidthPx(100)
	b.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETypeClick:

			appId := strings.TrimSpace(tb.Text())
			inviteCode := strings.TrimSpace(inviteTb.Text())
			acc := account.NewAccount(appId,inviteCode)

			if acc == nil {
				user.SetText("激活失败,请检查网络或者卡号")
				e.MarkDirty(user)
				return
			}

			if acc.IsActive() {
				user.SetText("软件已经激活:" + acc.TYPE())
				b.SetEnabled(false)
				e.MarkDirty(b)
				account.SaveAppId(appId)
				titleLabel.SetText(defaultShowString)
				e.MarkDirty(titleLabel)
				num ,code := account.AppAccount.InviteCode()
				l2.SetText(code)
				e.MarkDirty(l2)

				l3.SetText(fmt.Sprintf("已邀请 %d 人", num))
				e.MarkDirty(l3)


			} else {
				user.SetText("激活失败")
			}
			message.SetText(acc.Msg)
			timeStr.SetText(acc.TimeInfo())
			e.MarkDirty(user)
			e.MarkDirty(message)
			e.MarkDirty(timeStr)
		}
	}, gwu.ETypeClick)

	p.Add(b)
}

func buildInviteBottom(p gwu.Panel)  {
	buildInviteCodeUI(p)
	buildInviteUI(p)
}


func buildInviteCodeUI(p gwu.Panel)  {

	num ,code := account.AppAccount.InviteCode()

	p.AddVSpace(20)
	row := gwu.NewHorizontalPanel()

	l1 = gwu.NewLabel("你的邀请码是:")
	row .Add(l1)

	row.AddHSpace(10)
	l2= gwu.NewLabel(code)
	l2.Style().SetColor("red")
	l2.Style().SetFontSize("30")
	row .Add(l2)
	row.AddHSpace(10)
	l3 = gwu.NewLabel(fmt.Sprintf("已邀请 %d 人", num))
	row .Add(l3)

	p.Add(row)

}

func buildInviteUI(p gwu.Panel)  {

	p.AddVSpace(30)
	label := gwu.NewLabel("邀请机制")
	label.Style().SetColor("red")
	label.Style().SetFontSize("40")
	p.Add(label)
	p.AddVSpace(15)
	content := []string{
		"1. 购买卡号后可以获得邀请码",
		"2. 账户第一次激活使用邀请码才有效",
		"3. 使用邀请码激活，如果是月卡，双方各奖励10天，如果是年卡，双方各奖励30天",
		"4. 举例：张三买了一张月卡，使用了你的邀请码激活，则你的账户有效期增加10天，张三的有效期也增加10天",
		"5. 如果账户过期，邀请码也失效",
		"6. 公众号内分享的卡号没有邀请码",
	}
	for _, s := range content {
		p.Add(gwu.NewLabel(s))
		p.AddVSpace(10)
	}
}

