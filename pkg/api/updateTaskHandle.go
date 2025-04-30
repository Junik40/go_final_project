package api

import
(
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	datab "github.com/Junik40/go-final-project/pkg/db"
)



func updateTaskHandle (w http.ResponseWriter, r *http.Request) {
	var task datab.Tasks
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	defer r.Body.Close()
	Tittle := strings.ReplaceAll(task.Title, " ", "")
	if len(Tittle) == 0 {
		err = fmt.Errorf("tittle is empty")
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	err = check(&task)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	err = datab.UpdateTask(&task)
	if err != nil{

		writeJson(w,outEr{Error: err.Error()})	
		return
	}
	writeJson(w, map[string]string{})
	
}