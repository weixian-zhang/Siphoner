package main

import (
	"time"
)

type Config struct {
	SiphonFrequencySeconds int		`yaml:"SiphonFrequencySeconds"`
	TerminusSecretNamespace string	`yaml:"TerminusSecretNamespace"`
	SiphonKubeEvent bool			`yaml:"KubeEvent"`
	KeywordFilter []string			`yaml:"StdoutKeywordFilter"`
	NamespaceFilter []string 		`yaml:"IncludeNamespaces"`
	PodLabelsFilter []string 		`yaml:"IncludePodLabels"`
}

type PodInfo struct {
	Namespace string
	Name string
	Labels map[string]string
	ContainerNames []string
}

type LogResult struct {
	TimeGenerated time.Time
	Namespace string
	PodName string
	Container string
	Message string
}

type EventHubTerminusSecrets struct {

}