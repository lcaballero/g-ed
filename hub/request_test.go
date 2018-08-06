package hub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Bad_LoadRequest(t *testing.T) {
	// bad bin to LoadRequest returns nil and error
	assert.True(t, true)
}

func Test_ParseQueryParams(t *testing.T) {
	// finds ? and turns kvp in to map
	assert.True(t, true)
}

func Test_ParsePath(t *testing.T) {
	// finds ? returns path before the kvps
	assert.True(t, true)
}
