package api
import (

    "net/http"


)



func taskHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodPost:
		addTaskHandle(w, r)
	case http.MethodGet:
		getTasksHandle(w, r)
	default:

	

}
}