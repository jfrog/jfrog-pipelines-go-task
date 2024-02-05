package pipelines

import (
	"fmt"
	"github.com/jfrog/jfrog-pipelines-tasks-sdk-go/tasks"
	"net/url"
	"os"
	"strings"
)

// GetPipelinesURL returns pipelines URL of the instance by replacing "pielines_api_url"
// from v1 to pipelines
func GetPipelinesURL() string {
	url, err := url.Parse(GetStepURL())
	if err != nil {
		tasks.Warn(err.Error())
		return ""
	}
	hostname := strings.TrimPrefix(url.Hostname(), "")
	return fmt.Sprintf(url.Scheme + "://" + hostname + "/pipelines")
}

// GetStepURL fetches step URL where the taskis running
func GetStepURL() string {
	stepURL, err := tasks.GetVariable("step_url")
	if err != nil {
		tasks.Warn("Unable to fetch step url")
	}
	return stepURL
}

// PreparePipelineUIURL prepares pipelines UI URL using stepURL where the current task is
// getting executed. Fetches scheme, hostName from stepURl and project, pipelineName
// from generic integration created for monitoring pipelines.
func PreparePipelineUIURL(stepURL, project, pipelineName, branch string, runNumber int) string {
	url, err := url.Parse(stepURL)
	if err != nil {
		tasks.Warn(err.Error())
		return ""
	}
	hostname := strings.TrimPrefix(url.Hostname(), "")
	return fmt.Sprintf("%s://%s/ui/pipelines/myPipelines/%s/%s/%d?branch=%s", url.Scheme, hostname, project, pipelineName, runNumber, branch)
}

// GetProjectKey returns project key
func GetProjectKey() string {
	return os.Getenv("JFROG_CLI_BUILD_PROJECT")
}
