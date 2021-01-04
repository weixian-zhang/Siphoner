package main

import (
	
)

type Context struct {
	Config *Config
}

type Config struct {
	SiphonFrequencySeconds int		`yaml:"SiphonFrequencySeconds"`
	SiphonKubeEvent bool			`yaml:"KubeEvent"`
	StdoutKeywordFilter []string	`yaml:"StdoutKeywordFilter"` //1 or more keyword to filter
	StdoutRegexFilter []string		`yaml:"StdoutRegexFilter"` //or use regex expression. Regex takes priority
	NamespaceFilter []string 		`yaml:"IncludeNamespaces"`
	PodLabelsFilter []string 		`yaml:"IncludePodLabels"`
	
}

type PodInfo struct {
	Namespace string
	Name string
	Labels map[string]string
	ContainerNames []string
}