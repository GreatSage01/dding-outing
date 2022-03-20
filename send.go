package main

import (
	"app/dingtalk"

	"github.com/CodyGuo/glog"
)

func DtSend() {

	msg := []string{
		"### Link test",
		"---",
		"- <font color=#00ff00 size=6>红色文字</font>",
		"- content2",
	}

	dingToken := "595141398856ac5d2f614f0999029d4b5b7e74bcfa0faa2e8fa014e7b8d9f07c"
	dingSecret := "SEC29bf9cee13c396ca17c6a860aecd9fe462d558fed6ba23233d9c6d63f0d4796a"
	dt := dingtalk.InitDingTalkWithSecret(dingToken, dingSecret)
	if err := dt.SendMarkDownMessageBySlice("Markdown title", msg, dingtalk.WithAtAll()); err != nil {
		glog.Error(err)
	}
}
