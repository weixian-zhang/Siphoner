package main

import(
	"k8s.io/api/core/v1"
	"gopkg.in/yaml.v2"
)

func onConfigAdded(v1.ConfigMap) {
	
}

func onConfigUpdated(v1.ConfigMap) {

}

func getConfigMapFromList(cml *v1.ConfigMapList) {

}

func initConfig(cm v1.ConfigMap) (config Config, err error) {

	yamlStr := cm.Data["config"]

	config = Config{}

	err = yaml.Unmarshal([]byte(yamlStr), &config)
	Stdlog.Err(err)

	return config, err
}