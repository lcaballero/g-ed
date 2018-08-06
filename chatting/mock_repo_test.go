package chatting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type MockRepo map[string]interface{}

func NewMockRepo(kvp map[string]interface{}) MockRepo {
	return MockRepo(kvp)
}

func (m MockRepo) Set(key string, val interface{}) error {
	m[key] = val
	return nil
}

func (m MockRepo) Has(key string) bool {
	_, ok := m[key]
	return ok
}

func (m MockRepo) Get(key string) interface{} {
	val, ok := m[key]
	if ok {
		return val
	}
	return nil
}

func Test_Mock_Repo_Rooms(t *testing.T) {
	rooms := make(map[int]interface{})
	rooms[1] = Room{
		Occupants: []int{ 1 },
		Id: 1,
	}
	mr := make(MockRepo)
	mr.Set("rooms", rooms)

	assert.True(t, mr.Has("rooms"))
}

func Test_Mock_Set(t *testing.T) {
	mr := make(MockRepo)
	mr.Set("key1", "val1")
	_, ok := mr["key1"]
	assert.True(t, ok)
}

func Test_Mock_Has(t *testing.T) {
	mr := make(MockRepo)
	mr.Set("key1", "val1")
	assert.True(t, mr.Has("key1"))
}

func Test_Mock_Get(t *testing.T) {
	mr := make(MockRepo)
	mr.Set("key1", "val1")
	val := mr.Get("key1")
	assert.NotNil(t, val)
}
