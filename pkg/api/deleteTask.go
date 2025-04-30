package api

import (
	"net/http"

	datab "github.com/Junik40/go-final-project/pkg/db"
)




func deleteTaskHandle (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	defer r.Body.Close()
	err := datab.DeleteTask(id)
	if err != nil {
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	writeJson(w, map[string]string{}, http.StatusOK)
	
	




}