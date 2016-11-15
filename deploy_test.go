package werckerclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_deployTemplates(t *testing.T) {
	tests := []struct {
		Name     string
		Model    interface{}
		Expected string
	}{
		{"GetDeploys", &GetDeploysOptions{Owner: "foofoo", Name: "barbar"}, "/api/v3/applications/foofoo/barbar/deploys"},
		{"GetDeploys", &GetDeploysOptions{Owner: "foofoo", Name: "barbar", Status: "s", Limit: 123, Skip: 321, Sort: "so", Result: "r"}, "/api/v3/applications/foofoo/barbar/deploys?result=r&status=s&limit=123&skip=321&sort=so"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			uriTemplate, ok := deployTemplates[test.Name]
			require.True(t, ok, "uri template should exist")

			u, err := expandURL(uriTemplate, test.Model)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, u)
		})
	}
}
