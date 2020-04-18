package main

import (
	"log"
	"time"

	"github.com/y-shika/todo_golang_api/gateway"
)

func main() {
	db, err := gateway.NewMySQLClient()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("Succeed to connect MySQL")

	loadTodos(db)
}

func loadTodos(db *gateway.MySQLClient) {
	schema := `CREATE TABLE IF NOT EXISTS todos (
		id varchar(36) primary key,
		title text,
		active boolean,
		detail text,
		created_at datetime,
		updated_at datetime);`

	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}

	// TODO: fixtureのとあるレコードを消してもDBには反映されない
	//       ここのコードとDBの状態を常に同じにする手法がないか調査する
	todos := []struct {
		id     string
		title  string
		active bool
		detail string
	}{
		{
			id:     "1",
			title:  "Title_1",
			active: true,
			detail: "Detail_1",
		},
		{
			id:     "2",
			title:  "Title_2",
			active: true,
			detail: "Detail_2",
		},
		{
			id:     "3",
			title:  "Title_3",
			active: true,
			detail: "Detail_3",
		},
		{
			id:     "4",
			title:  "Title_4",
			active: true,
			detail: "Detail_4",
		},
		{
			id:     "5",
			title:  "Title_5",
			active: true,
			detail: "Detail_5",
		},
	}

	now := time.Now().Add(9 * time.Hour)

	for _, todo := range todos {
		ins, err := db.Prepare(`INSERT INTO todos (id, title, active, detail, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?) 
			ON DUPLICATE KEY UPDATE 
				title = VALUES(title),
				active = VALUES(active),
				detail = VALUES(detail),
				updated_at = VALUES(updated_at)`)
		if err != nil {
			panic(err)
		}

		_, err = ins.Exec(todo.id, todo.title, todo.active, todo.detail, now, now)
		if err != nil {
			panic(err)
		}
	}
}
