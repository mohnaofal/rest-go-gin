package mysql

import (
	"database/sql"
	"os"
	"time"
)

type mysqlConnection struct {
	mysql *sql.DB
}

type MySQLConnection interface {
	MySQL() *sql.DB
}

func (c *mysqlConnection) MySQL() *sql.DB {
	return c.mysql
}

func ConnectionDB() MySQLConnection {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("MYSQL_HOST")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mysqlConnection{
		mysql: db,
	}
}
