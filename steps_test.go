package werckerclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStepVersion(t *testing.T) {
	client := NewClient(&Config{})

	opts := &GetStepVersionOptions{
		Owner:   "wercker",
		Name:    "bash-template",
		Version: "*",
	}

	res, err := client.GetStepVersion(opts)
	assert.Nil(t, err)
	assert.Equal(t, res.Owner, "wercker")
}
