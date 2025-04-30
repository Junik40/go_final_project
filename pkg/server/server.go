package server

import (
	"fmt"
	"go1f/pkg/db"
	"net/http"
	"os"
    "go1f/pkg/api"

)
func Prepare() http.Handler{
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

    return http.FileServer(rootDirectory)
}



func Run(port string) {
    api.Init()
    fileServer := Prepare()
    http.Handle("/", fileServer)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Print(err)
    }

    
}
