package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoadSuccess(t *testing.T) {
	config := GetConfig()

	assert.Equal(t, ":6154", config.Server.Port)
	assert.Equal(t, "tcp", config.Server.Type)
}
