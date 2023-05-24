package config

type SMTP struct {
	Host     string `json:"host"`     // SMTP 服务器地址
	Port     int    `json:"port"`     // SMTP 服务器端口
	Username string `jspn:"username"` // 发件人邮箱用户名
	Password string `json:"password"` // 发件人邮箱密码
}
