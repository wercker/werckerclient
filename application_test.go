package werckerclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_applicationTemplates(t *testing.T) {
	tests := []struct {
		Name     string
		Model    interface{}
		Expected string
	}{
		{"GetApplication", &GetApplicationOptions{Owner: "foofoo", Name: "barbar"}, "/api/v3/applications/foofoo/barbar"},
		{"GetApplicationPipelines", &GetApplicationPipelinesOptions{Owner: "foofoo", Name: "barbar", Limit: "123", Skip: 321}, "/api/v3/applications/foofoo/barbar/pipelines?limit=123&skip=321"},
		{"GetApplicationPipelines", &GetApplicationPipelinesOptions{Owner: "foofoo", Name: "barbar"}, "/api/v3/applications/foofoo/barbar/pipelines"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			uriTemplate, ok := applicationTemplates[test.Name]
			require.True(t, ok, "uri template should exist")

			u, err := expandURL(uriTemplate, test.Model)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, u)
		})
	}
}
