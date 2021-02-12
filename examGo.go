package main

import (
	"fmt"
	"net/http"
	"time"
	"io / ioutil"
)

func main() {
	 
	http.HandleFunc("/", list)
	
	http.HandleFunc("/add", add)

	http.ListenAndServe(":8800", nil)
}


type Task struct {
	Description string
	Done bool
}
var tasks []Task



func list(rw http.ResponseWriter, _ *http.Request){
	for i,val := range tasks{
		if val.Done == false {
			chaine := fmt.Sprintf("ID: %d, task:%s",i , val.Description)
			bt := []byte (chaine)
			rw.WriteHeader(http.StatusOK)
			rw.Write(bt)
		}
	}



	type Welcome struct {
		Name string
		Time string
	}
	
func main() {
	
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	
	
	
	http.Handle("/myPath/", 
		http.StripPrefix("/static/",
				http.FileServer(http.Dir("static")))) 
				
	
	
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
	
			if name := r.FormValue("name"); name != "" {
		welcome.Name = name;
			}
			
	

	})

	func done(rw http.ResponseWriter, _ *http.Request){


	}
	
}
