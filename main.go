package main

import (
	"fmt"
	"log"
	"net/http"

	//	"os"
	"io/ioutil"
)

var serviceDiscovery string;
var dataURI string;
var deploymentName string;
func whoami(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"this is main \n")	
}

func getcount(w http.ResponseWriter, req *http.Request){
	resp, err := http.Get(dataUri + "/count")
	if err != nil {
		panic(err)	
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)	
	}
	fmt.Fprintf(w,dataURI)
}

func addcount(w http.ResponseWriter, req *http.Request){
	resp, err := http.Get(dataUri + "/addcount")
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
	http.HandleFunc("/whoami", whoami)
	http.HandleFunc("/count", getcount)
	http.HandleFunc("/addc`ount", addcount)
	log.Print("running server")
	log.Print(http.ListenAndServe("0.0.0.0:8080", nil))
}
