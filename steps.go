package werckerclient

import "github.com/jtacoma/uritemplates"

var stepTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(stepTemplates, "GetStepVersion", "/api/v2/steps{/owner,name,version}")
}

// StepService holds all run specific methods
type StepService interface {
	GetStepVersion(options *GetStepVersionOptions) (*Step, error)
}

// Ensure the client supports/adheres to the StepService interface
var _ StepService = (*Client)(nil)

// GetStepVersionOptions are the options associated with Client.GetStepVersion
type GetStepVersionOptions struct {
	Owner   string `map:"owner"`
	Name    string `map:"name"`
	Version string `map:"version"`
}

// GetStepVersion will retrieve a single Step
func (c *Client) GetStepVersion(options *GetStepVersionOptions) (*Step, error) {
	method := "GET"
	template := stepTemplates["GetStepVersion"]

	result := &Step{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
