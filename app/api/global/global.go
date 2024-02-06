package global

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"winter_test/app/api/global/config"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDB *sql.DB
	RDB     *redis.Client
)
