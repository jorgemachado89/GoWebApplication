package model

import (
	"database/sql"
)

// Hold pointer to DB object 
var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}