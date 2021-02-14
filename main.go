package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",Heart)
	if err:=http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal("Listen server failed.err: ",err)
	}
}


