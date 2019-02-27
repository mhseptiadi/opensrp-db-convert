package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	lib "./lib"
	repo "./repository"
)

const (
	DB_USER     = "opensrp_admin"
	DB_PASSWORD = "C4nT(T0ucH)Th1S"
	DB_NAME     = "opensrp"
	PORT        = "5430"
	RELPATH     = "./"
	SCHEMA      = "sid3"
)

func main() {

	lib.RELPATH = RELPATH
	lib.SCHEMA = SCHEMA

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=127.0.0.1 port=%s",
		DB_USER, DB_PASSWORD, DB_NAME, PORT)

	db, err := sql.Open("postgres", dbinfo)
	lib.CheckErr(err)
	defer db.Close()

	TambahAnc := repo.InitTambahAnc(db)
	TambahAnc.Run()

}
