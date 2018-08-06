package hub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Parse_Params(t *testing.T) {
	t.Log("should be able to parse an empty string")
	params := ParseQueryParams("/a/b/c?name=value")
	assert.Equal(t, "value", params["name"])
}

func Test_Parse_Path_With_Qs_Directories(t *testing.T) {
	t.Log("should be able to parse an empty string")
	path := ParsePath("/a/b/c?name=value")
	assert.Equal(t, "/a/b/c", path)
}

func Test_Parse_Path_With_Qs(t *testing.T) {
	t.Log("should be able to parse an empty string")
	path := ParsePath("directory?name=value")
	assert.Equal(t, "directory", path)
}

func Test_Parse_Path_Only_Query(t *testing.T) {
	t.Log("should be able to parse an empty string")
	path := ParsePath("?name=value")
	assert.Equal(t, "", path)
}

func Test_Parse_Path_Q(t *testing.T) {
	t.Log("should be able to parse an empty string")
	path := ParsePath("?")
	assert.Equal(t, "", path)
}

func Test_Parse_Path_Empty_String(t *testing.T) {
	t.Log("should be able to parse an empty string")
	path := ParsePath("")
	assert.Equal(t, "", path)
}
