package global

import (
	"file-manager/cores/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	CONFIG *config.Config
	VP     *viper.Viper
	LOG    *zap.Logger
)
