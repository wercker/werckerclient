package werckerclient

import (
	"fmt"

	"github.com/jtacoma/uritemplates"
)

var userTemplates = make(map[string]*uritemplates.UriTemplate)

func init() {
	addURITemplate(userTemplates, "GetUser", "/api/v2/profile")
	addURITemplate(userTemplates, "DeleteUser", "/api/v2/users")
}

type User struct {
	ID                string `json:"id"`
	FirstName         string `json:"name"`
	LastName          string `json:"lastName"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	GravatarHash      string `json:"gravatarHash"`
	Organization      string `json:"organisation"`
	URL               string `json:"url"`
	Twitter           string `json:"twitter"`
	HasGithubToken    bool   `json:"hasGithubToken"`
	HasBitbucketToken bool   `json:"hasBitbucketToken"`
}

func (u User) String() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type UserService interface {
	GetCurrentUser() (*User, error)
	DeleteUser() error
}

func (c *Client) GetCurrentUser() (*User, error) {
	method := "GET"
	template := userTemplates["GetUser"]

	result := &User{}
	err := c.Do(method, template, nil, nil, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) DeleteUser() error {
	method := "DELETE"
	template := userTemplates["DeleteUser"]

	err := c.Do(method, template, nil, nil, nil)
	return err
}
