package werckerclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_buildTemplates(t *testing.T) {
	tests := []struct {
		Name     string
		Model    interface{}
		Expected string
	}{
		{"GetBuilds", &GetBuildsOptions{Owner: "foofoo", Name: "barbar"}, "/api/v3/applications/foofoo/barbar/builds"},
		{"GetBuilds", &GetBuildsOptions{Owner: "foofoo", Name: "barbar", Commit: "c", Branch: "b", Status: "s", Limit: 123, Skip: 321, Sort: "so", Result: "r"}, "/api/v3/applications/foofoo/barbar/builds?commit=c&branch=b&status=s&limit=123&skip=321&sort=so&result=r"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			uriTemplate, ok := buildTemplates[test.Name]
			require.True(t, ok, "uri template should exist")

			u, err := expandURL(uriTemplate, test.Model)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, u)
		})
	}
}
