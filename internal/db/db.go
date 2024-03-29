package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const schema string = `
CREATE TABLE commands (
	id INTEGER PRIMARY KEY,
	command TEXT UNIQUE NOT NULL,
	description TEXT
)

CREATE TABLE queue (
	tele_handle TEXT PRIMARY KEY
)
`

func EstablishDBConnection() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
}