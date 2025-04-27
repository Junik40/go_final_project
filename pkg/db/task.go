package datab

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)



type Tasks struct {
    ID      string `json:"id"`
    Date    string `json:"date"`
    Title   string `json:"title"`
    Comment string `json:"comment"`
    Repeat  string `json:"repeat"`
}

func AddTask(task Tasks) (int64, error){
	var id int64
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	querry := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(querry, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	return id, err
}

func GetTasks(limit int) ([]Tasks, error) {
	var tasks []Tasks
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM scheduler LIMIT ? ORDER BY date ASC", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		fmt.Println(task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

