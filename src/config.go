package main

type SiphonerConfig struct {
	SiphonFrequencySeconds int	`yaml:"SiphonFrequencySeconds"`
	ConfigNamespace string		`yaml:"ConfigMapNamespace"`
	IncludeNamespaces []string 	`yaml:"IncludeNamespaces"`
	IncludePodLabels []string 	`yaml:"IncludePodLabels"`
}

func NewDefaultConfig() {
	
}