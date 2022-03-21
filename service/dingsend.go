package service

import (
	"strings"
	"zx/global"
	"zx/service/dingtalk"
)

func Dingsend(msg dingtalk.OutGoingModel) {
	session := strings.Split(msg.SessionWebhook, "=")[1]
	dt := dingtalk.InitDingTalkWithSession(session)

	// 字体及颜色
	dm := dingtalk.DingMap()
	dm.Set("颜色测试", dingtalk.H2)
	dm.Set("失败：$$ 同行不同色 $$", dingtalk.RED) // 双$分隔
	dm.Set("---", dingtalk.N)
	dm.Set("金色", dingtalk.GOLD)
	dm.Set("成功", dingtalk.GREEN)
	dm.Set("警告", dingtalk.BLUE)
	dm.Set("普通文字", dingtalk.N)
	if err := dt.SendMarkDownMessageBySlice("color test", dm.Slice()); err != nil {
		global.ZX_LOG.Error(err.Error())
	}
}
