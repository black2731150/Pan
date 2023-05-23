package config

type SMTP struct {
	Host     string // SMTP 服务器地址
	Port     int    // SMTP 服务器端口
	Username string // 发件人邮箱用户名
	Password string // 发件人邮箱密码
}
