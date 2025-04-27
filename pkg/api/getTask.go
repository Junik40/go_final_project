package api

import (
	"fmt"
	"net/http"

	datab "github.com/Junik40/go-final-project/pkg/db"
)

func getTasksHandle(w http.ResponseWriter, r *http.Request){
	var tasks []datab.Tasks
	tasks, err := datab.GetTasks(50)
	if err != nil {
		writeJson(w, err)
		return
	}
	fmt.Print(tasks)
	writeJson(w, tasks)
}