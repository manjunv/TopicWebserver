// Copyright 2017-2018 IBM Inc. All rights reserved.

package main

import (
	"fmt"
	"net/http"
	"os"
)

var path = "/Users/mgovin739/go/src/server/form.html"


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) { return }
}

func httphandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    
    switch r.Method {

    case "GET":     
         http.ServeFile(w, r, "form.html")

    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from Topic ! TOPIC = %v\n", r.PostForm)
        name := r.FormValue("name")
        address := r.FormValue("address")
        fmt.Fprintf(w, "Topic = %s\n", name)
        fmt.Fprintf(w, "Topic Content = %s\n",address)

    case "DELETE":
	deleteFile()
    default:
        fmt.Fprintf(w, "nil methods are not supported.")
    }
}

 
func main() {
    http.HandleFunc("/",httphandler)
    fmt.Printf("Starting Webserver for Topics/Blogs handling and POST/GET/DELETE..\n")
    if err := http.ListenAndServe("localhost:8080", nil); err != nil {
        check(err)
    }
}
