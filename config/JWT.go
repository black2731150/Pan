package config

type JWTconfig struct {
	Secret string `yaml:"secret"`
	Issure string `yaml:"issure"`
	Expire int    `yaml:"expire"`
}
