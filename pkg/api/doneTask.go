package api

import (
	"net/http"  
	datab "github.com/Junik40/go-final-project/pkg/db"
	"time"
)

func doneHandler (w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == ""{
		writeJson(w,outEr{Error: "id is empty"}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	task,err := datab.GetTask(id)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	if task.Repeat == ""{
		err = datab.DeleteTask(id)
		if err != nil{
			writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
			return
		}
		writeJson(w, map[string]string{}, http.StatusOK)
		return
	}
	now := time.Now()
	task.Date,err = NextDate(now, task.Date, task.Repeat)
	
	if err != nil{
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	err = datab.UpdateTask(&task)
	if err != nil{

		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)	
		return
	}
	writeJson(w, map[string]string{}, http.StatusOK)
	


	



}