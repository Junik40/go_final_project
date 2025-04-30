package api

import (
	"net/http"
)



func taskHandler(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case http.MethodGet:
		getTaskHandle(w, r)
	case http.MethodPost:
		addTaskHandle(w, r)	
	case http.MethodPut:
		updateTaskHandle(w, r)
	case http.MethodDelete:
		deleteTaskHandle(w, r)
	default:	
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	

}
}