package main

type TerminusLogAppender interface {
	Init(config Config)
	AppendLog(log LogResult)
}