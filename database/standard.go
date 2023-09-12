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

func SelectMyData(id int) MyData {
	con := connect()
	defer con.Close()

	q := "select * from mydata where id = $1"
	rs := con.QueryRow(q, id)

	var md MyData
	err := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if err != nil {
		log.Fatalln("row fetch error", err)
	}

	return md
}

func CreateMyData(md *MyData) {
	con := connect()
	defer con.Close()

	q := "insert into mydata (name, mail, age) values ($1, $2, $3)"
	con.Exec(q, md.Name, md.Mail, md.Age)
}

func UpdateMyData(md *MyData) {
	con := connect()
	defer con.Close()

	q := "update mydata set name=$1, mail=$2, age=$3 where Name='PGUMA'"
	con.Exec(q, md.Name, md.Mail, md.Age)
}

func DeleteMyData(name string) {
	con := connect()
	defer con.Close()

	q := "delete from mydata where name=$1"
	con.Exec(q, name)
}

func connect() *sql.DB {
	// TODO 環境変数等から接続文字列を取得する
	con, err := sql.Open("postgres", "postgres://pguma:password@localhost:5432/dev?sslmode=disable")
	if err != nil {
		log.Fatalln("failed to connect DB", err)
	}
	return con
}
