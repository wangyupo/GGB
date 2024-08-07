package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/wangyupo/GGB/config"
	"github.com/wangyupo/GGB/utils/timer"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GGB_DB     *gorm.DB
	GGB_REDIS  redis.UniversalClient
	GGB_CONFIG config.Server
	GGB_VP     *viper.Viper
	GGB_LOG    *zap.Logger
	GGB_Timer  timer.Timer = timer.NewTimerTask()
	GGB_Trans  ut.Translator
)
