package main

import(
	"k8s.io/api/core/v1"
)

func onConfigAdded(v1.ConfigMap) {
	
}

func onConfigUpdated(v1.ConfigMap) {

}

func getConfigMapFromList(cml *v1.ConfigMapList) {

}

func initConfig(cm v1.ConfigMap) (config Config) {
	return
}