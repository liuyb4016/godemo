package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.PATH=%q\n",r.URL.Path)
}
