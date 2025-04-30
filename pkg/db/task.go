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

func GetTasks(limit int) ([]*Tasks, error) {
	var tasks []Tasks
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT * FROM scheduler ORDER BY date ASC LIMIT %d", limit)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close()
	var tasksp []*Tasks
	for i := range tasks{tasksp = append(tasksp, &tasks[i])}
	return tasksp, nil
}
func GetTask(id string) (Tasks, error) {
	var task Tasks
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return task, err
	}
	defer db.Close()
	check := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM scheduler WHERE id = %s)", id)
	var exists bool
	err = db.QueryRow(check).Scan(&exists)
	if err != nil {
		return task, err
	}
	if !exists {
		return task, fmt.Errorf("incorrect id")
	}
	query := fmt.Sprintf("SELECT * FROM scheduler WHERE ID = %s", id)
	row := db.QueryRow(query)
	row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)

	return task, nil
	
}
func UpdateTask(task *Tasks) error {
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE scheduler SET date = '%s', title = '%s', comment = '%s', repeat = '%s' WHERE id = %s", task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	res, err := db.Exec(query)
	if err != nil {
		return err
	}
    count, err := res.RowsAffected()
    if err != nil {
        return err
    }
    if count == 0 {
        return fmt.Errorf(`incorrect id for updating task`)
    }
    return nil
}
func DeleteTask(id string) error {
	db, err := sql.Open("sqlite", "./pkg/db/scheduler.db")
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("DELETE FROM scheduler WHERE id = %s", id)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}




	return nil



}