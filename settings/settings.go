package settings

import "github.com/spf13/viper"

var (
	ConfViper *viper.Viper
	Config    config
)

func SetupConfig(configFile string) {
	ConfViper = viper.New()
	ConfViper.SetConfigFile(configFile)
	ConfViper.SetConfigType("yml")
	if err := ConfViper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := ConfViper.UnmarshalKey("settings", &Config); err != nil {
		panic(err)
	}
}
