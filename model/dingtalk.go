package model

type OutGoingModel struct {
	AtUsers []struct {
		DingtalkID string `json:"dingtalkId"`
	} `json:"atUsers"`
	ChatbotUserID             string `json:"chatbotUserId"`
	ConversationID            string `json:"conversationId"`
	ConversationTitle         string `json:"conversationTitle"`
	ConversationType          string `json:"conversationType"`
	CreateAt                  int64  `json:"createAt"`
	IsAdmin                   bool   `json:"isAdmin"`
	IsInAtList                bool   `json:"isInAtList"`
	MsgID                     string `json:"msgId"`
	Msgtype                   string `json:"msgtype"`
	SceneGroupCode            string `json:"sceneGroupCode"`
	SenderID                  string `json:"senderId"`
	SenderNick                string `json:"senderNick"`
	SessionWebhook            string `json:"sessionWebhook"`
	SessionWebhookExpiredTime int64  `json:"sessionWebhookExpiredTime"`
	Text                      struct {
		Content string `json:"content"`
	} `json:"text"`
}

