package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
)

//Implement singleton
func GetMySqlConnection() (*sqlx.DB, error) {
	Db, err := sqlx.Connect("mysql", "root:ons_vg@tcp(127.0.0.1:3306)/BOOK_MY_SHOW")
	if err != nil {
		return nil, errors.Wrap(err, "connection to mysql database failed")
	}
	log.Println("connection to mysql database established successfully")
	return Db, err
}
