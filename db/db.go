package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New(user string, password string, database string) (*sql.DB, error) {
	return sql.Open("mysql", user + ":" + password + "@/" + database + "?charset=utf8&parseTime=True&loc=Local")
}