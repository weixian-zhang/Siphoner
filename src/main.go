package main

import (
	"github.com/weixian-zhang/Siphoner/kube"
)

var Stdlog Logger


func main() {

	Stdlog = &StdLog{}
	Stdlog.Init()
	
	initKubeClientSet()

	// ns := []string{"dev", "default"}
	// podL := make(map[string]string)
	// podL["run"] = "helloworld"
	// pods, _ := GetPodsByFilteredNamespaces(ns, podL)
	// getPodLogs(pods)

	initConfigFromConfigMapList()

}