package config

import "time"

type Config struct {
	ZapConfig      ZapConfig      `mapstructure:"zap_config" json:"zap_config"`
	DatabaseConfig DatabaseConfig `mapstructure:"database" json:"database"`
	ServerConfig   ServerConfig   `mapstructure:"server" json:"server"`
}

type ServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Mode string `mapstructure:"mode" json:"mode"`
}

type ZapConfig struct {
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

type DatabaseConfig struct {
	MysqlConfig `mapstructure:"mysql" json:"mysql"`
	RedisConfig `mapstructure:"redis" json:"redis"`
}

type MysqlConfig struct {
	Addr            string        `mapstructure:"addr" json:"addr"`
	Port            string        `mapstructure:"port" json:"port"`
	DB              string        `mapstructure:"db" json:"db"`
	Username        string        `mapstructure:"username" json:"username"`
	Password        string        `mapstructure:"password" json:"password"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifeTime" json:"connMaxLifeTime"`
	ConnMaxIdleTime time.Duration `mapstructure:"connMaxIdleTime" json:"connMaxIdleTime"`
	MaxIdleConns    int           `mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int           `mapstructure:"maxOpenConns" json:"maxOpenConns"`
	charset         int           `mapstructure:"charset" json:"charset"`
	Place           int           `mapstructure:"place" json:"place"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Port     string `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
	Host     string `mapstructure:"host" json:"host"`
}

type User struct {
	Name     string
	Age      int
	Location string
}
