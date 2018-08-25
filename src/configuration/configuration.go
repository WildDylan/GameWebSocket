package configuration

import "github.com/jinzhu/configor"

var Config = struct {
	AppName             string
	AppDesc             string

	Redis struct {
		Host            string
		Port            string
		Password        string
		DB              int
	}

	Contacts struct {
		Author          string
		Email           string
	}
}{}

func LoadConfig() {
	configor.Load(&Config, "src/configuration/config.yml")
}
