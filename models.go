package werckerclient

import "time"

// Build is a detailed api representation
type Build struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Application *ApplicationSummary `json:"application"`
	Branch      string              `json:"branch"`
	CommitHash  string              `json:"commitHash"`
	CreatedAt   time.Time           `json:"createdAt"`
	EnvVars     []EnvVar            `json:"envVars"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Message     string              `json:"message"`
	Progress    int                 `json:"progress"`
	Result      string              `json:"result"`
	StartedAt   time.Time           `json:"startedAt"`
	Status      string              `json:"status"`
}

// BuildSummary is a summary api representation
type BuildSummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Branch     string    `json:"branch"`
	CommitHash string    `json:"commitHash"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Message    string    `json:"message"`
	Progress   int       `json:"progress"`
	Result     string    `json:"result"`
	StartedAt  time.Time `json:"startedAt"`
	Status     string    `json:"status"`
}

// Run is a detailed api representation
type Run struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Application *ApplicationSummary `json:"application"`
	Branch      string              `json:"branch"`
	Tag         string              `json:"tag"`
	CommitHash  string              `json:"commitHash"`
	CreatedAt   time.Time           `json:"createdAt"`
	EnvVars     []EnvVar            `json:"envVars"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Message     string              `json:"message"`
	Progress    int                 `json:"progress"`
	Result      string              `json:"result"`
	StartedAt   time.Time           `json:"startedAt"`
	SourceRun   *RunSummary         `json:"sourceRun"`
	Pipeline    *PipelineSummary    `json:"pipeline"`
	Status      string              `json:"status"`
}

// RunSummary is a summary api representation
type RunSummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Branch     string    `json:"branch"`
	Tag        string    `json:"tag"`
	CommitHash string    `json:"commitHash"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Message    string    `json:"message"`
	Progress   int       `json:"progress"`
	Result     string    `json:"result"`
	StartedAt  time.Time `json:"startedAt"`
	Status     string    `json:"status"`
}

// RunStep is a detailed api representation
type RunStep struct {
	ID           string              `json:"id"`
	URL          string              `json:"url"`
	ArtifactsURL string              `json:"artifactsUrl"`
	CreatedAt    time.Time           `json:"createdAt"`
	FinishedAt   time.Time           `json:"finishedAt"`
	Message      string              `json:"message"`
	LogURL       string              `json:"logUrl"`
	Order        int                 `json:"order"`
	Application  *ApplicationSummary `json:"project"`
	Phase        string              `json:"phase"`
	Result       string              `json:"result"`
	Run          *RunSummary         `json:"run"`
	StartedAt    time.Time           `json:"startedAt"`
	Status       string              `json:"status"`
	Step         string              `json:"step"`
}

// Deploy is a detailed api representation
type Deploy struct {
	ID          string              `json:"id"`
	URL         string              `json:"url"`
	Status      string              `json:"status"`
	Result      string              `json:"result"`
	CreatedAt   time.Time           `json:"createdAt"`
	FinishedAt  time.Time           `json:"finishedAt"`
	Progress    int                 `json:"progress"`
	Application *ApplicationSummary `json:"application"`
	Build       *BuildSummary       `json:"build"`
}

// DeploySummary is a summary api representation
type DeploySummary struct {
	ID         string    `json:"id"`
	URL        string    `json:"url"`
	Status     string    `json:"status"`
	Result     string    `json:"result"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
	Progress   int       `json:"progress"`
}

// EnvVar represents a environment variable key value pair
type EnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func EnvVarFromMap(envmap map[string]string) []EnvVar {
	envvars := []EnvVar{}
	for k, v := range envmap {
		envvars = append(envvars, EnvVar{Key: k, Value: v})
	}
	return envvars
}

// SCM is a detailed source control manager api representation
type SCM struct {
	Type       string `json:"type"`
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
}

// Application is a detailed api representation
type Application struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	Builds    string       `json:"builds"`
	Deploys   string       `json:"deploys"`
	SCM       *SCM         `json:"scm"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

// Pipeline is a detailed api representation
type Pipeline struct {
	ID                   string    `json:"id"`
	Url                  string    `json:"url"`
	CreatedAt            time.Time `json:"createdAt"`
	Name                 string    `json:"name"`
	Permissions          string    `json:"permissions"`
	PipelineName         string    `json:"pipelineName"`
	SetSCMProviderStatus bool      `json:"setScmProviderStatus"`
	Type                 string    `json:"type"`
}

// ApplicationSummary is a summary api representation
type ApplicationSummary struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Name      string       `json:"name"`
	Owner     *UnifiedUser `json:"owner"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Privacy   string       `json:"privacy"`
	Stack     int          `json:"stack"`
}

// PipelineSummary is a summary api representation
type PipelineSummary struct {
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"createdAt"`
	Name                 string    `json:"name"`
	Permissions          string    `json:"permissions"`
	PipelineName         string    `json:"pipelineName"`
	SetSCMProviderStatus bool      `json:"setScmProviderStatus"`
	Type                 string    `json:"type"`
}

// Step
/*
{
  "deploy": "57a49b50ec801e0100d7b8d9",
  "applicationOwner": {
    "name": "",
    "username": "wercker",
    "gravatar": "33a5bfbcf8a2b90f40e849b6f1fa5eeb"
  },
  "application": "546bae9be0193716350004fa",
  "createdOn": "2016-08-05T13:58:22.688Z",
  "name": "bash-template",
  "fullname": "wercker/bash-template",
  "owner": "wercker",
  "version": "2016.218.1354",
  "description": "templates files using environment variables",
  "license": null,
  "werckerUrl": null,
  "codeUrl": null,
  "tarballUrl": "https://s3.amazonaws.com/wercker-production-steps/5cbd017a-6ddd-40a8-b26b-d1987d084fe9",
  "tarballSize": 1009,
  "shasum": "d031a72edc34b5cddaac494b2f2c22e9a505b040d90b213beebff472a23185f0",
  "readMe": "Templates files containing .template using environment variables\n\nexample.template\n\n  Hello ${PWD}\n\n\nafter using this step you will have a\n\nexample\n\n  Hello /pipeline/source\n\n\n",
  "keywords": [
    "bash",
    "template"
  ],
  "packageExclude": null,
  "type": null,
  "reportFilename": null,
  "reportType": null,
  "applicationOwnerGravatar": "33a5bfbcf8a2b90f40e849b6f1fa5eeb",
  "properties": {
    "output": {
      "default": "",
      "required": false,
      "type": "string"
    },
    "input": {
      "default": "",
      "required": false,
      "type": "string"
    }
  },
  "main": "run.sh",
  "viewCount": 0,
  "packageVersion": "1",
  "releasedBy": {
    "username": "termie",
    "gravatar": "75d28dd33caf573c352d5afa937f4476"
  }
}
*/

type Step struct {
	Deploy           string     `json:"deploy"`
	ApplicationOwner *StepUser  `json:"applicationOwner"`
	Application      string     `json:application"`
	CreatedOn        *time.Time `json:"createdOn"`
	Name             string     `json:"name"`
	Fullname         string     `json:"fullname"`
	Owner            string     `json:"owner"`
	Version          string     `json:"version"`
	Description      string     `json:"description"`
	License          string     `json:"license"`
	WerckerURL       string     `json:"werckerUrl"`
	CodeURL          string     `json:"codeUrl"`
	TarballURL       string     `json:"tarballUrl"`
	TarballSize      int        `json:"tarballSize"`
	Shasum           string     `json:"shasum"`
	Readme           string     `json:"readMe"`
	Keywords         []string   `json:"keywords"`

	// No idea the format of these / deprecated
	// PackageExclude string `json:"packageExclude"`
	// Type string `json:"type"`
	// ReportFilename `json:"reportFilename"`
	// ReportType `json:"reportType"`

	// Duplicate data
	// ApplicationOwnerGravatar string `json:"applicationOwnerGravatar"`

	Properties     map[string]*StepProperty `json:"properties"`
	Main           string                   `json:"main"`
	ViewCount      int                      `json:"viewCount"`
	PackageVersion string                   `json:"packageVersion"`
	ReleasedBy     *StepUser                `json:"releasedBy"`
}

// StepUser is the user representation used by the step detail
type StepUser struct {
	Name     string `json:"name"`
	Gravatar string `json:"gravatar"`
	Username string `json:"username"`
}

type StepProperty struct {
	Default  string `json:"default"`
	Required bool   `json:"required"`
	Type     string `json:"type"`
}

// Token is a detailed api representation
type Token struct {
	ID             string     `json:"id"`
	URL            string     `json:"url"`
	Name           string     `json:"name"`
	Token          string     `json:"token"`
	HashedToken    string     `json:"hashedToken"`
	LastCharacters string     `json:"lastCharacters"`
	CreatedAt      *time.Time `json:"createdAt"`
	LastUsedAt     *time.Time `json:"lastUsedAt"`
}

// TokenSummary is a summary api representation
type TokenSummary struct {
	ID             string     `json:"id"`
	URL            string     `json:"url"`
	Name           string     `json:"name"`
	HashedToken    string     `json:"hashedToken"`
	LastCharacters string     `json:"lastCharacters"`
	CreatedAt      *time.Time `json:"createdAt"`
	LastUsedAt     *time.Time `json:"lastUsedAt"`
}

// UnifiedUser is a flexible user representation. Not all fields have to be set
type UnifiedUser struct {
	Type   string             `json:"type"`
	Name   string             `json:"name"`
	Avatar *UnifiedUserAvatar `json:"avatar"`
	UserID string             `json:"userId"`
	Meta   *UnifiedUserMeta   `json:"meta"`
}

// UnifiedUserAvatar is the avatar property of the UnifiedUser
type UnifiedUserAvatar struct {
	Gravatar string `json:"gravatar"`
}

// UnifiedUserMeta is the meta property of the UnifiedUser
type UnifiedUserMeta struct {
	Username        string `json:"username"`
	WerckerEmployee bool   `json:"werckerEmployee"`
}
