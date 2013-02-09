package main

import (
	_ "github.com/bmizerany/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/eaigner/hood"
)

var (
	hd *hood.Hood
)

type Room struct {

}

type User struct {

}

func init() {
	//db, err := sql.Open("postgres", "user=pg_spyke_user dbname=spyke sslmode=verify-full")
	hd, err := hood.Open("sqlite3", "db=./_database.db")

	hd.CreateTableIfNotExists(&Room{})
	hd.CreateTableIfNotExists(&User{})
}