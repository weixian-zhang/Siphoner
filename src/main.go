package main

var Stdlog Logger


func main() {

	Stdlog = &StdLog{}
	Stdlog.Init()

	err :=  GetPodLogs("default", "nginx", "")
	Stdlog.Err(err)
}