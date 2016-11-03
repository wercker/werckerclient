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
