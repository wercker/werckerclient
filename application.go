package werckerclient

import "github.com/jtacoma/uritemplates"

// applicationTemplates contains all UriTemplates indexed by name.
var applicationTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(applicationTemplates, "GetApplication", "/api/v3/applications{/applicationName}")
	addURITemplate(applicationTemplates, "GetApplicationPipelines", "/api/v3/applications{/applicationName}/pipelines{?limit,skip}")
}

// GetApplicationOptions are the options associated with Client.GetApplication
type GetApplicationOptions struct {
	ApplicationName string `map:"applicationName"`
}

// GetApplication will retrieve a single Application
func (c *Client) GetApplication(options *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := applicationTemplates["GetApplication"]

	result := &Application{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ApplicationService interface {
	GetApplication(*GetApplicationOptions) (*Application, error)
}

// GetApplicationsOptions are the options associated with Client.GetApplications
type GetApplicationsOptions struct {
	Limit string `map:"limit,omitempty"`
	Skip  int    `map:"skip,omitempty"`
	Sort  int    `map:"sort,omitempty"`
	Stack string `map:"stack,omitempty"`
}

// GetApplications will retrieve a list of Applications
func (c *Client) GetApplications(options *GetApplicationsOptions) ([]Application, error) {
	method := "GET"
	template := applicationTemplates["GetApplications"]

	result := []Application{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type GetApplicationPipelinesOptions struct {
	// Required
	ApplicationName string `map:"applicationName"`

	// Optional
	Limit string `map:"limit,omitempty"`
	Skip  int    `map:"skip,omitempty"`
}

// GetApplicationPipelines will retrieve a list of of an Application's pipelines
func (c *Client) GetApplicationPipelines(options *GetApplicationPipelinesOptions) ([]PipelineSummary, error) {
	method := "GET"
	template := applicationTemplates["GetApplicationPipelines"]

	result := []PipelineSummary{}
	err := c.Do(method, template, options, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
