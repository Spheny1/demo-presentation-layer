package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var serviceDiscovery string;
var dataURI string;
var deploymentName string;
func whoami(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"this is a feature branch\n")	
}

func getcount(w http.ResponseWriter, req *http.Request){
	resp, err := http.Get(dataURI + "/count")
	if err != nil {
		panic(err)	
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)	
	}
	fmt.Fprintf(w,string(body))
}

func addcount(w http.ResponseWriter, req *http.Request){
	_ , err := http.Get(dataURI + "/addcount")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w,"Incremented \n")	
	_ , err = http.Get(dataURI + "/addcount")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w,"Incremented \n")	
}

func main(){
	serviceDiscovery=os.Getenv("SERVICE_DISCOVERY_URI")
	deploymentName = os.Getenv("DEPLOYMENT_NAME")
	resp, err := http.Get(serviceDiscovery + "/resolve?name=" + deploymentName)
	if err != nil {
		panic(err)	
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)	
	}
	dataURI = string(body)
	if len(dataURI) < 5 {
		panic("restart service due to empty resp from service-discovery")
	}
	http.HandleFunc("/whoami", whoami)
	http.HandleFunc("/count", getcount)
	http.HandleFunc("/addcount", addcount)
	log.Print("running server")
	log.Print(http.ListenAndServe("0.0.0.0:8080", nil))
}
