package main

import (
	"fmt"

	"github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"
)

func main() {
	tasks.Info("Starting task ...")
	userName := tasks.GetInput("user")
	message := fmt.Sprintf("Hello %s from pipelines tasks 😊", userName)
	// Set greeting message as task output
	tasks.SetOutput("message", message)
}