package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root123@/registry?charset=utf8")
	checkErr(err)

	// stmt, err := db.Prepare("insert into user (username, email, password, realname) values (?, ?, ?, ?)")

	// res, err := stmt.Exec("stone", "stone@gmail.com", "123456", "zhang")

	// checkErr(err)

	// id, err := res.LastInsertId()

	// checkErr(err)

	// fmt.Println(id)

	rows, err := db.Query("select username, email, password, realname from user")
	checkErr(err)

	for rows.Next() {
		var username string
		var email string
		var password string
		var realname string

		err = rows.Scan(&username, &email, &password, &realname)
		checkErr(err)
		fmt.Println("==========================================")
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Println(realname)
	}

	createUserAndProject(db)

	db.Close()

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createUserAndProject(db *sql.DB) error {
	tx, _ := db.Begin()
	defer tx.Rollback()

	if _, err := tx.Exec("insert into user (username, email, password, realname) values ('apple', 'apple@example.com', '123456', 'common user')"); err != nil {
		return nil
	}

	if _, err := tx.Exec("insert into project (owner_id, name) values (1, 'orange_library')"); err != nil {
		return nil
	}

	err := tx.Commit()
	return err
}
