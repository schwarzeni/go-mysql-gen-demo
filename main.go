package main

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// RESET QUERY CACHE
// set profiling=1;
// show profiles;
func main() {
	db, err := sql.Open("mysql", "root:ilovetos@tcp(10.248.166.131:3306)/larget_data?charset=utf8")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	count := 500000
	// count := 4000
	maxIdx := count * 2
	group := 5000
	rand.Seed(time.Now().Unix())
	for count > 0 {
		num := group
		if count < group {
			num = count
		}
		sqlStmt := "INSERT INTO authors (idx, name) VALUES "
		data := make([]string, num)
		for i := range data {
			data[i] = "(" + strconv.Itoa(rand.Intn(maxIdx)) + ", " + "'nzy'" + ")"
		}
		sqlStmt += strings.Join(data, " , ")
		if _, err := db.Exec(sqlStmt); err != nil {
			panic(err)
		}
		count -= group
	}

	// insert, err := db.Prepare("INSERT INTO authors (idx) VALUES (?)")
	// if err != nil {
	// 	panic(err)
	// }
	// begin, err := db.Begin()
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < count; i++ {
	// 	_, err := begin.Stmt(insert).Exec(rand.Intn(count * 10))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// err = begin.Commit()
	// if err != nil {
	// 	panic(err)
	// }
}

type Author struct {
	ID  int
	Idx int
}
