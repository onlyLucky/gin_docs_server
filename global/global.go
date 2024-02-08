package global

import (
	"gin_docs_server/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *logrus.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)


/* 
// 扩展 log
var (
	Config *config.Config
	Log *LogServer
)

type LogServer struct {
	*logrus.Logger
	ServiceName string
} 
// main.js
global.Log = new(global.LogServer)
global.Log.Logger = core.InitLogger()
global.Log.ServiceName = "logger"
*/

