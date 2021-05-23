package db

import (
	"log"

	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func InitBuntDB() {
	var err error

	// Open the data.db file. It will be created if it doesn't exist.
	db, err = buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateIndex("name", "*", buntdb.IndexJSON("name"))
}

func Close() {
	db.Close()
}

func GetDB() *buntdb.DB {
	return db
}
