# {{ .Name }}

This JFrog pipelines task performs this action.
 
this task performs below actions

- Does some awesome work
- Installation of a popular package
- Performs some repetitive tasks
- Sends notification

### What's New

`Add whats latest in the release`
- Added to support caching

### Prerequisites

`Add pre requisites for running this task`

This task requires jfrog-cli, JFrog Xray

## Usage

Give some sample configuration for running 
**Basic:**

```yaml
- task: jfrog/{{ .Name }}@v0.0.1
  repository: pipelines-tasks-virtual
  id: pipe_{{ .Name }}
  input:
    {{ range .Inputs }}
    {{ .Name }}: value
    {{ end }}
```

### Input Variables

| Name                        | Required | Default                               | Description                     |
|-----------------------------|----------|---------------------------------------|---------------------------------|
{{ range .Inputs }}
| {{ .Name }}                  | {{ .Required }} | {{ .DefaultValue }}            | {{ .Description }}              |
{{ end }}

### Exported Environment Variables

#### Output Variables

{{ range .Outputs }}
- {{ .Name }}
{{ end }}

### How does it work?

Give a simple description how Task functions

## License

Mention the License for the Task

## Release Notes

The release notes are available [here](RELEASE.md).

## Troubleshooting

- Add any known errors and what might be the reason for the error.

## Related Tasks

- Mention all the relatable tasks which might help.