package app

import "github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"

// AUTO Generated
type Inputs struct {
	{{- range .Inputs }}
	{{ .Name }}	string
	{{- end}}
}

// AUTO Generated
func Run() error {
	i := &Inputs{}
	{{- range .Inputs }}
	i.{{ .Name }} = tasks.GetInput("{{ .Name }}")
	{{- end}}
	err := i.validateInputs()
	if err != nil {
		return err
	}
	return nil
}

// AUTO Generated
// validateInputs performs additional validations required for task
func (i *Inputs) validateInputs() error {
	// Perform any extra validations required for task
	
	return nil
}

// AUTO Generated
// SetTaskOutputs sets tasks outputs
func SetTaskOutputs(output... string) {
	{{- range $i, $a := .Outputs }}
	tasks.SetOutput("{{ .Name }}", output[{{ $i }}])
	{{- end}}
}

// AUTO Generated
// GetIntegrationValue Read integration from jfrog pipelines and get values
func GetIntegrationValue(i Inputs) (string, error) {
	// Pass the input name expected to get integration
	integration, err := tasks.GetIntegration()
	if err != nil {
		return "", err
	}
	// Visit https://jfrog.com/help/r/jfrog-pipelines-documentation/pipelines-integrations to find out
	// the string to be sent to integration.GetValue() to read integration values
	return integration.GetValue()
}

// AUTO Generated
// GetResourceValue
func GetResourceValue(i Inputs) (string, error) {
	// Pass the input name expected to get resource
	resource, err := tasks.GetResource()
	if err != nil {
		return "", err
	}
	// returns ResourcePath
	resourcePath := resource.ResourcePath
	// returns resource type
	resourceType := resource.ResourceType
}
