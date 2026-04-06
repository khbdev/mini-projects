package main

import "net/http"


type User struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}


func Handler(w http.ResponseWriter, r *http.Request){
  switch r.Method {
  case http.MethodPost:
	
  }
}