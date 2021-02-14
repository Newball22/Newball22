package main

import "net/http"

func Heart(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("hello world"))
}