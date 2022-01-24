package demo

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestQuery(t *testing.T) {
	//db, err := sql.Open("mysql", "dev:123654@tcp(47.98.248.153:9036)/demo")
	db, err := sql.Open("mysql", "dev:123654@tcp(127.0.0.1:3336)/demo")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	rows, err := db.Query("SELECT * FROM category WHERE category_id > ?", 0)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	var (
		id   int
		name string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		println(">>cate:", id, name)
	}
}
