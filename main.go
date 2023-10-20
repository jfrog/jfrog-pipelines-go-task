package main

import "github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"

func main() {
	tasks.Info("Starting task ...")
	status := "success"
	// Set greeting message as task output
	tasks.SetOutput("status", status)
}
