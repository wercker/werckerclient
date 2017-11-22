package werckerclient

import (
	"github.com/jtacoma/uritemplates"
)

var organizationTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(organizationTemplates, "DeleteUser", "/api/v2/organizations{/username}")
}

type OrganizationService interface {
	DeleteOrganization(options *DeleteOrganizationOptions) error
}

// DeleteOrganizationOptions are the options associated with Client.DeleteOrganization
type DeleteOrganizationOptions struct {
	Username string `map:"username"`
}

func (c *Client) DeleteOrganization(options *DeleteOrganizationOptions) error {
	method := "DELETE"
	template := organizationTemplates["DeleteOrganization"]

	err := c.Do(method, template, options, nil, nil)
	return err
}
