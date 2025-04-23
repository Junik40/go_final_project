package server

import (
	"fmt"
	"net/http"
    "os"
    "go1f/pkg/db"
)

func Run(port string) {
    

    
    dbFile := "./pkg/db/scheduler.db"
    _, err := os.Stat(dbFile)

    var install bool
    if err != nil {
        install = true
    }
    if install {
        datab.Create_Table(dbFile)
    }
    rootDirectory := http.Dir("web")

    fileServer := http.FileServer(rootDirectory)


    http.Handle("/", fileServer)
    err = http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Print(err)
    }
// если install равен true, после открытия БД требуется выполнить 
// sql-запрос с CREATE TABLE и CREATE INDEX

}
