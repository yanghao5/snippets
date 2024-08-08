package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 打开数据库连接，如果数据库文件不存在则创建它
	db, err := sql.Open("sqlite3", "./version.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建一个表
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS img_version (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		ver TEXT NOT NULL
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO img_version(ver) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("sss/myalpine:01")
	if err != nil {
		log.Fatal(err)
	}

	// query data
	rows, err := db.Query("SELECT id, ver FROM img_version")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var ver string
		err = rows.Scan(&id, &ver)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %s \n", id, ver)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// 查询最大 id
	var maxID int
	err = db.QueryRow("SELECT MAX(id) FROM img_version").Scan(&maxID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(maxID)
}
