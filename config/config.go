package config

import "github.com/mohnaofal/rest-go-gin/config/mysql"

type Config struct {
	mysqlDB mysql.MySQLConnection
}

func LoadConfig() *Config {
	cfg := new(Config)

	cfg.InitMySQLDB()

	return cfg
}

func (c *Config) InitMySQLDB() {
	c.mysqlDB = mysql.ConnectionDB()
}

func (c *Config) MySQLDB() mysql.MySQLConnection {
	return c.mysqlDB
}
