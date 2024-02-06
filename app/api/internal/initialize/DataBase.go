package initialize

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"time"
	"winter_test/app/api/global"
)

func SetupDataBase() {
	setupMysql()
	setupRedis()
}

func setupMysql() {
	config := global.Config.DatabaseConfig.MysqlConfig

	db, err := sql.Open("mysql", config.Addr+":"+config.Port)
	if err != nil {
		global.Logger.Fatal("open mysql failed," + err.Error())
	}

	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		global.Logger.Fatal("connect to mysql failed," + err.Error())
	}
	global.MysqlDB = db

	global.Logger.Info("init mysql success")
}

func setupRedis() {
	config := global.Config.DatabaseConfig.RedisConfig

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	})
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("connect to redis falied")
	}
	global.RDB = rdb

	global.Logger.Info("init redis success")

}
