package global

import (
	"file-manager/cores/config"
	"github.com/spf13/viper"
)

var (
	CONFIG *config.Config
	VP     *viper.Viper
)
