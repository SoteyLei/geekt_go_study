package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func wrapSqlErro(conn *sql.DB) (string, error) {
	var username, password string
	err1 := conn.QueryRow("select userName from user where userName = ?", "abcd").Scan(&username, &password)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			return nil, errors.Wrap(err1, "not found")
		} else {
			log.Fatal(err1)
		}
	}
	fmt.Println(username)
}

func main() {

	conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/testtest?charset=utf8")
	if err != nil {
		fmt.Println("链接失败")
	}
	res, err := conn.Exec("create table user(userName VARCHAR(30),passwd VARCHAR(40))")
	fmt.Println("create table result=", res, err)

	res, err = conn.Exec("insert user(userName,passwd) values (?,?)", "aaa", "bbb")

	wrapSqlErro(conn)

	defer conn.Close()
}
