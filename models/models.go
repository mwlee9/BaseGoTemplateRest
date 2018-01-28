package models

import (
	"database/sql"
	"fmt"
)

// Database
func InitDatabase() *sql.DB {
	conn := "postgres://swyreijf:hlR3e6UqP7YEsy6nq_BIChyRE8SPINoP@nutty-custard-apple.db.elephantsql.com:5432/swyreijf"
	db, err := sql.Open("postgres", conn)

	checkErr(err)

	return db

}

// CreateTable ...
func CreateTable(db *sql.DB) {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Animals (id TEXT, name TEXT, species TEXT);")

	checkErr(err)

	err1 := db.Close()

	checkErr(err1)

	defer db.Close()

}

// GetAllTasks ...
func GetAllTasks() *sql.Rows {
	db := InitDatabase()

	rows, err := db.Query("SELECT * FROM animals;")

	checkErr(err)

	defer db.Close()

	return rows
}

// GetOneTask ...
func GetOneTask(params string) *sql.Rows {
	db := InitDatabase()
	rows, err := db.Query("SELECT * FROM animals WHERE id =" + `'` + params + `'`)

	checkErr(err)

	defer db.Close()

	return rows
}

// DeleteOneTask ...
func DeleteOneTask(params string) {
	db := InitDatabase()
	stmt, err := db.Prepare("DELETE FROM animals WHERE id = $1;")

	checkErr(err)

	res, err1 := stmt.Exec(params)

	checkErr(err1)

	fmt.Println(res)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
