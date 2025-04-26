package api
import (
    "net/http"
	"go1f/pkg/db"

)

func addTaskHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodPost:
		var task =  datab.Tasks

		task := datab.Tasks {
			Date: r.FormValue("date"),
			Title: r.FormValue("title"),
			Comment: r.FormValue("comment"),
			Repeat: r.FormValue("repeat"),
		} 




	}

}