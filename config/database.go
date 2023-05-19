package config

//数据库基础配置
type Database struct {
	Driver   string   `yaml:"driver"`   //数据库驱动
	Host     string   `yaml:"host"`     //主机ip
	Port     int      `yaml:"port"`     //端口
	Username string   `yaml:"username"` //数据库用户名
	Password string   `yaml:"password"` //密码
	Name     string   `yaml:"name"`     //数据库名字
	Options  []string `yaml:"options"`  //配置选项
}
