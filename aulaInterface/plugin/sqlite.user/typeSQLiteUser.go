package main

import (
	"database/sql"
)

var User SQLiteUser

type SQLiteUser struct {
	Database *sql.DB
}
