package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}
