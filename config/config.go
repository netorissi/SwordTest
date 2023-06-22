package config

import (
	"github.com/netorissi/SwordTest/shared"
	"github.com/spf13/viper"
)

var Global *ConfigGlobal

func LoadConfig() {
	var (
		v = viper.New()
		f = "./config/development.json"
	)

	v.AutomaticEnv()
	v.SetConfigFile(f)

	if err := v.ReadInConfig(); err != nil {
		shared.Logger.DPanic(err)
	}

	if err := v.Unmarshal(&Global); err != nil {
		shared.Logger.DPanic(err)
	}

	shared.Logger.Info("Config loaded.")
}
