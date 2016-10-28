package werckerclient

import "github.com/jtacoma/uritemplates"

// runTemplates contains all UriTemplates indexed by name.
var runStepTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(runTemplates, "GetRunStep", "/api/v3/runsteps{/runStepId}")
}

// RunStepService holds all runStep specific methods
type RunStepService interface {
	GetRunStep(options *GetRunStepOptions) (*RunStep, error)
}

// Ensure the client supports/adheres to the RunService interface
var _ RunStepService = (*Client)(nil)

// GetRunStepOptions are the options associated with Client.GetRunStep
type GetRunStepOptions struct {
	RunStepID string `map:"runStepId"`
}

// GetRunStep will retrieve a single RunStep
func (c *Client) GetRunStep(options *GetRunStepOptions) (*RunStep, error) {
	method := "GET"
	template := runTemplates["GetRunStep"]

	result := &RunStep{}
	err := c.Do(method, template, options, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
