package api

import (
	"strconv"
	"strings"
	"time"
	"net/http"
	"fmt"
)


func AfterNow(date, now time.Time) bool {
	date = date.Truncate(24 * time.Hour)
	now = now.Truncate(24 * time.Hour)
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
		if interval >400 || interval < 1{
			err = fmt.Errorf("выход за промежуток")
			return "", err
		}

		for {
			date = date.AddDate(0, 0, interval)

			if AfterNow(date, now) {
				break
			}
		}
		return date.Format("20060102"), nil
	case  "y":
		for {
			date = date.AddDate(1, 0, 0)
			if AfterNow(date, now) {
				break
			}

		}
		return date.Format("20060102"), nil

	default:
		err = fmt.Errorf("ошибочный формат повторения")
		return "", err 
	}
		

}

func NextDayHandler(w http.ResponseWriter, r *http.Request){
	prenow := r.URL.Query().Get("now")
	now, err := time.Parse("20060102", prenow)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	dstart := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

    date, err := NextDate(now, dstart, repeat)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()})
		return
	}
	date1, err := strconv.Atoi(date)
	if err != nil{
		writeJson(w,outEr{Error: err.Error()})
		return
	}

	writeJson(w, date1)
    
}
