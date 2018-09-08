package server

import "github.com/jmoiron/sqlx"

var (
	Db *sqlx.DB
)

func Init(db *sqlx.DB) (err error) {
	Db = db
	return
}
