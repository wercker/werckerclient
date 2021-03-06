package werckerclient

import (
	"errors"

	"github.com/jtacoma/uritemplates"
)

// applicationTemplates contains all UriTemplates indexed by name.
var applicationTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(applicationTemplates, "GetApplication", "/api/v3/applications{/owner,name}")
	addURITemplate(applicationTemplates, "GetApplicationPipelines", "/api/v3/applications{/owner,name}/pipelines{?limit,skip}")
}

// GetApplicationOptions are the options associated with Client.GetApplication
type GetApplicationOptions struct {
	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`

	// ApplicationName will override Owner and Name
	ApplicationName string `map:"applicationName"`
}

// GetApplication will retrieve a single Application
func (c *Client) GetApplication(o *GetApplicationOptions) (*Application, error) {
	method := "GET"
	template := applicationTemplates["GetApplication"]

	if o.ApplicationName != "" {
		if owner, name, ok := parseApplicationName(o.ApplicationName); ok {
			o.Owner = owner
			o.Name = name
		} else {
			return nil, errors.New("Unable to parse ApplicationName")
		}
	}

	result := &Application{}
	err := c.Do(method, template, o, nil, result)
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
	Name  string `map:"name"`
	Owner string `map:"owner"`

	// Optional
	Limit string `map:"limit,omitempty"`
	Skip  int    `map:"skip,omitempty"`

	// ApplicationName will override Owner and Name
	ApplicationName string `map:"-"`
}

// GetApplicationPipelines will retrieve a list of of an Application's pipelines
func (c *Client) GetApplicationPipelines(o *GetApplicationPipelinesOptions) ([]PipelineSummary, error) {
	method := "GET"
	template := applicationTemplates["GetApplicationPipelines"]

	if o.ApplicationName != "" {
		if owner, name, ok := parseApplicationName(o.ApplicationName); ok {
			o.Owner = owner
			o.Name = name
		} else {
			return nil, errors.New("Unable to parse ApplicationName")
		}
	}

	result := []PipelineSummary{}
	err := c.Do(method, template, o, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
