package main

import (
	"net/http"
	"fmt"
)

func main()  {

    http.HandleFunc("/", index_handler)
    http.HandleFunc("/login", login_handler)
    http.ListenAndServe(":8000",nil)
}

func index_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Go Web App..! Yayyy....")
}

func login_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Go Web App..! Login....")
}