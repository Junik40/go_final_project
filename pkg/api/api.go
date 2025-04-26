package api
import (
    "net/http"
)

func Init() {
    http.HandleFunc("/api/nextdate", NextDayHandler)
	http.HandleFunc("/api/task", addTaskHandler)

}