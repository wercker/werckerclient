package werckerclient

import (
	"github.com/jtacoma/uritemplates"
)

var organizationTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(organizationTemplates, "DeleteOrganization", "/api/v2/organizations{/username}")
}

// OrganizationService holds all organization specific methods
type OrganizationService interface {
	DeleteOrganization(options *DeleteOrganizationOptions) error
}

// DeleteOrganizationOptions are the options associated with Client.DeleteOrganization
type DeleteOrganizationOptions struct {
	Username string `map:"username"`
}

// DeleteOrganization removes an organization
func (c *Client) DeleteOrganization(options *DeleteOrganizationOptions) error {
	method := "DELETE"
	template := organizationTemplates["DeleteOrganization"]

	err := c.Do(method, template, options, nil, nil)
	return err
}
