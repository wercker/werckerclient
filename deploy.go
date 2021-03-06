package werckerclient

import (
	"errors"

	"github.com/jtacoma/uritemplates"
)

// deployTemplates contains all UriTemplates indexed by name.
var deployTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(deployTemplates, "GetDeploy", "/api/v3/deploys{/deployId}")
	addURITemplate(deployTemplates, "GetDeploys", "/api/v3/applications{/owner,name}/deploys{?buildId,result,stack,status,limit,skip,sort}")
}

// GetDeployOptions are the options associated with Client.GetDeploy
type GetDeployOptions struct {
	DeployID string `map:"deployId"`
}

// GetDeploy will retrieve a single Deploy
func (c *Client) GetDeploy(options *GetDeployOptions) (*Deploy, error) {
	method := "GET"
	template := deployTemplates["GetDeploy"]

	result := &Deploy{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeploysOptions are the options associated with Client.GetDeploys.
type GetDeploysOptions struct {
	// Required
	Owner string `map:"owner"`
	Name  string `map:"name"`

	// Optional
	BuildID string `map:"buildId,omitempty"`
	Limit   int    `map:"limit,omitempty"`
	Result  string `map:"result,omitempty"`
	Skip    int    `map:"skip,omitempty"`
	Sort    string `map:"sort,omitempty"`
	Stack   string `map:"stack,omitempty"`
	Status  string `map:"status,omitempty"`

	// ApplicationName will override Owner and Name
	ApplicationName string `map:"-"`
}

// GetDeploys fetches all deploys for a certain application and optional filters.
func (c *Client) GetDeploys(o *GetDeploysOptions) ([]*DeploySummary, error) {
	method := "GET"
	template := deployTemplates["GetDeploys"]

	if o.ApplicationName != "" {
		if owner, name, ok := parseApplicationName(o.ApplicationName); ok {
			o.Owner = owner
			o.Name = name
		} else {
			return nil, errors.New("Unable to parse ApplicationName")
		}
	}

	result := []*DeploySummary{}
	err := c.Do(method, template, o, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
