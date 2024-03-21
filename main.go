package main

import (
	"log"
	"fmt"
	"net/http"
)

func whoami(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"this is main \n")	
}

func getcount(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"this is main \n")	
}

func addcount(w http.ResponseWriter, req *http.Request){
fmt.Fprintf(w,"this is main \n")	
}

func main(){
	http.HandleFunc("/whoami", whoami)
	http.HandleFunc("/count", getcount)
	http.HandleFunc("/addcoount", addcount)
	log.Print("running server")
	log.Print(http.ListenAndServe("0.0.0.0:8080", nil))
}
