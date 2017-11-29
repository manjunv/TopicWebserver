// Copyright 2017-2018 IBM Inc. All rights reserved.

package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"path/filepath"
)

var path = "$HOME/form"


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

func deleteFile(file string) {
	var err = os.Remove(file)
	if isError(err) { return }
}


func updateTopic(file string,req *http.Request) { 
	var write []byte
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
        }
	outfile, err := os.Open(file)
    		check(err)
    	if err != nil {
       		panic(err)
    	}
    	defer outfile.Close()
	body, err := ioutil.ReadAll(req.Body)
	write = append(write, []byte(body)...)
	write = append(write, body...)

	//write it to a file
	err = ioutil.WriteFile(file,write, 0644)
	if err != nil {
    		panic(err)
	}
}
 
func httphandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    file := filepath.Base(path)
    switch r.Method {

    case "GET":     
         http.ServeFile(w, r,file)

    case "POST":
	if _, err := os.Stat(file); os.IsNotExist(err) {
		updateTopic(file,r)
      }
	updateTopic(file,r)
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
    case "DELETE":
	deleteFile(file)
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
