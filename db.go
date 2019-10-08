package database

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func GetDatabase()(db * sql.DB, err  error){

	db , err =  sql.Open("sqlite3", "./api/database/test.db")
	return

}