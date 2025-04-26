package api

import (
	"strconv"
	"strings"
	"time"
	"net/http"
)


func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == ""{
		return "", nil
	}
	date,err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}

	input := strings.Split(repeat, " ")
	form := input[0]

	switch form {
	case "d":
		if len(input) < 2 {
			return "", nil
		}
		interval,err := strconv.Atoi(input[1])
		if err != nil {
			return "", err
		}
		if interval >400{
			return "",nil
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	case  "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	
	default:
		return "", nil
	}
		

}

func NextDayHandler(w http.ResponseWriter, r *http.Request){
	prenow := r.FormValue("now")
	now, err := time.Parse("20060102", prenow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dstart := r.FormValue("date")
	repeat := r.FormValue("repeat")

    date, err := NextDate(now, dstart, repeat)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(date))
    
}
