package database_standard

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type MyData struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m *MyData) Str() string {
	e, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return "{ error }"
	}
	return string(e)
}

func SelectMyDataAll() {

	con := connect()
	defer con.Close()

	q := "select * from mydata"
	rs, err := con.Query(q)
	if err != nil {
		log.Fatalln("query error", err)
	}

	for rs.Next() {
		var md MyData
		err := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if err != nil {
			log.Fatalln("row fetch error", err)
		}
		fmt.Println(md.Str())
	}
}

func connect() *sql.DB {
	// TODO 環境変数等から接続文字列を取得する
	con, err := sql.Open("postgres", "postgres://pguma:password@localhost:5432/dev?sslmode=disable")
	if err != nil {
		log.Fatalln("failed to connect DB", err)
	}
	return con
}
