package config

type Database struct {
	Driver   string   `yaml:"driver"`
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Name     string   `yaml:"name"`
	Options  []string `yaml:"options"`
}
