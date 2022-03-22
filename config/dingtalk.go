package config

type Dingtalk struct {
	AgentId   int    `mapstructure:"agentId" json:"agentId" yaml:"agentId"`
	AppKey    string `mapstructure:"appKey" json:"appKey" yaml:"appKey"`
	AppSecret string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
