package config

//服务配置
type Configuration struct {
	Pan       Pan       `yaml:"pan"`       //pan基础配置
	Database  Database  `yaml:"database"`  //数据库基础配置
	SMTP      SMTP      `yaml:"smtp"`      //邮件服务基础配置
	JWTconfig JWTconfig `yaml:"JWTconfig"` //JWT配置信息
}
