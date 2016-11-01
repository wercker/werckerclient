package werckerclient

import "github.com/jtacoma/uritemplates"

// runTemplates contains all UriTemplates indexed by name.
var runTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(runTemplates, "CreateRun", "/api/v3/runs")
	addURITemplate(runTemplates, "GetRun", "/api/v3/runs{/runId}")
	addURITemplate(runTemplates, "GetRuns", "/api/v3/runs{?applicationId,pipelineId,limit,skip,sort,status,result,branch,pipelineId,commit,sourceRun,author}")
}

// RunService holds all run specific methods
type RunService interface {
	GetRun(options *GetRunOptions) (*Run, error)
	GetRuns(options *GetRunsOptions) ([]*RunSummary, error)
	CreateRun(options *CreateRunOptions) (*Run, error)
	CreateChainRun(options *CreateChainRunOptions) (*Run, error)
}

// Ensure the client supports/adheres to the RunService interface
var _ RunService = (*Client)(nil)

// GetRunOptions are the options associated with Client.GetRun
type GetRunOptions struct {
	RunID string `map:"runId"`
}

// GetRun will retrieve a single Run
func (c *Client) GetRun(options *GetRunOptions) (*Run, error) {
	method := "GET"
	template := runTemplates["GetRun"]

	result := &Run{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetRunsOptions are the options associated with Client.GetRuns.
// Options defined here: http://devcenter.wercker.com/api/endpoints/runs.html#get-all-runs
type GetRunsOptions struct {
	// Required (one or the other)
	ApplicationID string `map:"owner"`
	PipelineID    string `map:"name"`

	// Optional
	Limit     int    `map:"limit,omitempty"`
	Skip      int    `map:"skip,omitempty"`
	Sort      string `map:"sort,omitempty"`
	Status    string `map:"status,omitempty"`
	Result    string `map:"result,omitempty"`
	Branch    string `map:"branch,omitempty"`
	Commit    string `map:"commit,omitempty"`
	SourceRun string `map:"sourceRun,omitempty"`
	Author    string `map:"author,omitempty"`
}

// GetRuns fetches all runs for a certain application and optional filters.
func (c *Client) GetRuns(options *GetRunsOptions) ([]*RunSummary, error) {
	method := "GET"
	template := runTemplates["GetRuns"]

	result := []*RunSummary{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateRunOptions are the options associated with Client.CreateRun.
type CreateRunOptions struct {
	// Required
	PipelineID string `json:"pipelineId,omitempty"`

	// Optional
	Branch     string   `json:"branch,omitempty"`
	CommitHash string   `json:"commitHash,omitempty"`
	Message    string   `json:"message,omitempty"`
	EnvVars    []EnvVar `json:"envVars,omitempty"`
}

// CreateRun will trigger a new run.
func (c *Client) CreateRun(options *CreateRunOptions) (*Run, error) {
	method := "POST"
	template := runTemplates["CreateRun"]

	result := &Run{}
	err := c.Do(method, template, options, options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateChainRunOptions are the options associated with Client.CreateChainRun.
type CreateChainRunOptions struct {
	// Required
	SourceRunID string `json:"sourceRunId,omitempty"`
	TargetID    string `json:"targetId,omitempty"`

	Message string   `json:"message,omitempty"`
	EnvVars []EnvVar `json:"envVars,omitempty"`
}

// CreateChainRun will trigger a new run.
func (c *Client) CreateChainRun(options *CreateChainRunOptions) (*Run, error) {
	method := "POST"
	template := runTemplates["CreateRun"]

	result := &Run{}
	err := c.Do(method, template, options, options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
