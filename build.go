package werckerclient

import (
	"errors"

	"github.com/jtacoma/uritemplates"
)

// buildTemplates contains all UriTemplates indexed by name.
var buildTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(buildTemplates, "CreateBuild", "/api/v3/builds")
	addURITemplate(buildTemplates, "GetBuild", "/api/v3/builds{/buildId}")
	addURITemplate(buildTemplates, "GetBuilds", "/api/v3/applications{/owner,name}/builds{?commit,branch,status,limit,skip,sort,result}")
}

// GetBuildOptions are the options associated with Client.GetBuild
type GetBuildOptions struct {
	BuildID string `map:"buildId"`
}

// GetBuild will retrieve a single Build
func (c *Client) GetBuild(options *GetBuildOptions) (*Build, error) {
	method := "GET"
	template := buildTemplates["GetBuild"]

	result := &Build{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetBuildsOptions are the options associated with Client.GetBuilds.
type GetBuildsOptions struct {
	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`

	// Optional
	Branch string `map:"branch,omitempty"`
	Commit string `map:"commit,omitempty"`
	Limit  int    `map:"limit,omitempty"`
	Result string `map:"result,omitempty"`
	Skip   int    `map:"skip,omitempty"`
	Sort   string `map:"sort,omitempty"`
	Stack  string `map:"stack,omitempty"`
	Status string `map:"status,omitempty"`

	// ApplicationName will override Owner and Name
	ApplicationName string `map:"-"`
}

// GetBuilds fetches all builds for a certain application and optional filters.
func (c *Client) GetBuilds(o *GetBuildsOptions) ([]*BuildSummary, error) {
	method := "GET"
	template := buildTemplates["GetBuilds"]

	if o.ApplicationName != "" {
		if owner, name, ok := parseApplicationName(o.ApplicationName); ok {
			o.Owner = owner
			o.Name = name
		} else {
			return nil, errors.New("Unable to parse ApplicationName")
		}
	}

	result := []*BuildSummary{}
	err := c.Do(method, template, o, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateBuildOptions are the options associated with Client.CreateBuild.
type CreateBuildOptions struct {
	// Required
	ApplicationID string `json:"applicationId,omitempty"`

	// Optional
	Branch     string   `json:"branch,omitempty"`
	CommitHash string   `json:"commitHash,omitempty"`
	Message    string   `json:"message,omitempty"`
	EnvVars    []EnvVar `json:"envVars,omitempty"`
}

// CreateBuild will trigger a new build.
func (c *Client) CreateBuild(options *CreateBuildOptions) (*Build, error) {
	method := "POST"
	template := buildTemplates["CreateBuild"]

	result := &Build{}
	err := c.Do(method, template, options, options, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
