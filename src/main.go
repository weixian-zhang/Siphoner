package main

var Stdlog Logger


func main() {

	Stdlog = &StdLog{}
	Stdlog.Init()
	
	initKubeClientSet()

	ns := []string{"dev", "default"}
	podL := make(map[string]string)
	podL["run"] = "helloworld"

	pods, _ := GetPodsByFilteredNamespaces(ns, podL)

	getPodLogs(pods)
	
	//Stdlog.Err(err)
	

	
	// Stdlog.Err(err)
}