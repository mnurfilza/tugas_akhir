package main

import (
	"database/sql"
	"log"
	"net/http"
	"project/belajargolang/tugasakhir"
	"project/belajargolang/tugasakhir/lib"
)

var user, password, host, port, dbname string

func init() {
	user = "root"
	password = ""
	host = "127.0.0.1"
	port = "3306"
	dbname = "sistemkampus"
}

func main() {
	db, err := lib.Connect(user, password, host, port, dbname)
	if err != nil {
		db, err = InitDb()
		if err != nil {
			return
		}
	}
	defer db.Close()
	tugasakhir.RegisDB(db)

	http.HandleFunc("/api/ss/", tugasakhir.SS)

	log.Println("localhost : 8089")
	http.ListenAndServe(":8089", nil)
}

func InitDb() (*sql.DB, error) {
	db, err := lib.Connect(user, password, host, port, "")
	if err != nil {
		return nil, err
	}
	if err := lib.DropDB(db, dbname); err != nil {
		return nil, err
	}

	if err := lib.CreateDatabase(db, dbname); err != nil {
		return nil, err
	}

	db, err = lib.Connect(user, password, host, port, dbname)
	if err != nil {
		return nil, err
	}

	if err = lib.Use(db, dbname); err != nil {
		return nil, err
	}

	if err = lib.CreateTable(db, tugasakhir.TblMahasiswa); err != nil {
		return nil, err
	}

	return db, err
}
