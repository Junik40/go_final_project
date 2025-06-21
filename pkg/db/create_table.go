package datab

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)


func Create_Table(name string) {

	schema1 := `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date DATE NOT NULL,
		title VARCHAR(256) NOT NULL,
		comment TEXT NOT NULL DEFAULT "",
		repeat VARCHAR(128)
	);
	`
	schema2 := `
	CREATE INDEX date On scheduler(date); 
	`
	db, err := sql.Open("sqlite", name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec(schema1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(schema2)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}
