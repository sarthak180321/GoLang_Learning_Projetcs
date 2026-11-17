package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
func helper(writer http.ResponseWriter,request *http.Request){
	writer.Header().Set("content-type","application/json")
	writer.WriteHeader(http.StatusOK) //status code 200
	res:=map[string]any{
		"status":"ok",
		"message":"json encoded sucessfully",
		"time":time.Now().UTC(),
	}
	_=json.NewEncoder(writer).Encode(res)
	//usin NewENcode returns You a new encode which writes in the writer
}
func Test() {
    http.HandleFunc("/ok", helper)

    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        fmt.Println("error occured while starting server", err.Error())
    }
}