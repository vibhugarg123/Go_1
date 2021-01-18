package config

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func GetMySqlConnection() (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("mysql", "root:root@(127.0.0.0:3306)/mysql")
	if err != nil {
		log.Fatalln(err)
	}
	return
}
