package api

import (
	"encoding/json"
	"net/http"

)

func writeJson(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if _,ok := data.(error); ok {
        w.WriteHeader(http.StatusBadRequest)
        data = map[string]string{"error": data.(error).Error()}
        err:= json.NewEncoder(w).Encode(data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else {
        w.WriteHeader(http.StatusOK)
        data = map[string]any{"id": data}
        err:= json.NewEncoder(w).Encode(data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }

}


func Init() {
    http.HandleFunc("/api/nextdate", NextDayHandler)
	http.HandleFunc("/api/task", taskHandler)

}