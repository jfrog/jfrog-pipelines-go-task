package main

import (
	"jfrog-pipelines-go-task/app"
	"os"

	"github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"
)

func main() {
	tasks.Info("Starting {{ .Name }} ...")
	err := app.Run()
	if err != nil {
		Exit(err)
	}
}

func Exit(err error) {
	tasks.Error(err.Error())
	os.Exit(1)
}
