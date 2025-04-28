package api

import (
	"encoding/json"
	"net/http"

)

func writeJson(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json, charset=utf-8")
    json.NewEncoder(w).Encode(data)

}


func Init() {
    http.HandleFunc("/api/nextdate", NextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
    http.HandleFunc("/api/tasks", getTasksHandler )
    http.HandleFunc("/api/task/done",doneHandler)
    
} 