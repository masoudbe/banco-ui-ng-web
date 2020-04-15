package main

import (
    "fmt"
    "net/http"
)
import "io/ioutil"
import "encoding/json"

type ActionInfo struct {
    Command string `json:"command"`
    Data    string `json:"data"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(fileName string) string  {
    data, err := ioutil.ReadFile(fileName)
    check(err)
    return string(data)
}

func execute(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    for k, v := range r.URL.Query() {
        fmt.Printf("%s: %s\n", k, v)
    }

    var params = r.URL.Path

    b, _ := ioutil.ReadAll(r.Body)

    var actionInfo ActionInfo
    json.Unmarshal(b, &actionInfo)

    var result = ""

    if params == "/api/commandexecution" {
        result = readFile("commandexecution.json")
    }

    if params == "/api/CommandLinkGroups" {
        result = readFile("commandlinkgroups.json")
    }
    if params == "/api/CommandLinks" {
        result = readFile("commandlinks.json")
    }

    json.NewEncoder(w).Encode(result)
}

func main() {
    //router := mux.NewRouter()
    //router.HandleFunc("/execute", execute).Methods("GET")

    http.HandleFunc("/api/CommandLinkGroups", execute)
    http.HandleFunc("/api/CommandLinks", execute)
    http.HandleFunc("/api/commandexecution", execute)
    //http.Handle("/", http.FileServer(http.Dir("/bancong")))

    fmt.Println("server is running at port 8888")

    http.ListenAndServe(":8888", nil)

    fmt.Scanln()
}
