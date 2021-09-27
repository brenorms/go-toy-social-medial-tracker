package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

func GetDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:r00t123@(localhost:3306)/socialmediatracker") //localhost, not available
	return db, err
}
