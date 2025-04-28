package api

import (
	"fmt"
	"net/http"

	datab "github.com/Junik40/go-final-project/pkg/db"
)

func getTaskHandle (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		writeJson(w,outEr{Error: "id is empty"})
		return
	}
	defer r.Body.Close()
	task, err := datab.GetTask(id)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	writeJson(w, task)
	fmt.Println(task)
	
}