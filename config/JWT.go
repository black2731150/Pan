package config

type JWTconfig struct {
	Secret string `json:"secret"`
	Issure string `json:"issure"`
	Expire int    `json:"expire"`
}
