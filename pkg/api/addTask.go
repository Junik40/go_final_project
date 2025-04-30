package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	datab "github.com/Junik40/go-final-project/pkg/db"
	"strings"
)

type outEr struct {
	Error string `json:"error"`
}

type outId struct {
	Id string `json:"id"`
}


func check(task *datab.Tasks) error {
	now := time.Now()
	if task.Date == ""{
		task.Date = now.Format("20060102")
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil{
		return err
	}
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil{
		return err
	}
	
	if AfterNow(now, t) {
        if len(task.Repeat) == 0 {
            task.Date = now.Format("20060102")
			
        } else {
            task.Date = next
        }
    } 
	return nil
}



func addTaskHandle(w http.ResponseWriter, r *http.Request) {
	var task datab.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	Tittle := strings.ReplaceAll(task.Title, " ", "")
	if len(Tittle) == 0 {
		err = fmt.Errorf("tittle is empty")
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	err = check(&task)
	if err != nil {
		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	id,err := datab.AddTask(task)
	if err != nil {

		writeJson(w,outEr{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	writeJson(w, outId{Id: strconv.FormatInt(id, 10)}, http.StatusOK)
	

}