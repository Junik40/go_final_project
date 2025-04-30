package api

import (
	"net/http"

	datab "github.com/Junik40/go-final-project/pkg/db"
)

type outTasks struct {
	Tasks []*datab.Tasks `json:"tasks"`
}
func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := datab.GetTasks(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tasks == nil {
		tasks = []*datab.Tasks{}
	}
	writeJson(w, outTasks{Tasks: tasks}, http.StatusOK)
}	
