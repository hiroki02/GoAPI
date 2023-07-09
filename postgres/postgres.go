package postgres

import (
	"database/sql"
	"log"
	"time"
)

type TaskList []Task

type Task struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "user=User password='' host=localhost port=5432 dbname=lesson sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func ListTask() (TaskList, error) {
	var t Task
	var tl TaskList
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&t.ID, &t.Title, &t.UpdatedAt, &t.CreatedAt)
		tl = append(tl, t)
	}
	return tl, err

}
