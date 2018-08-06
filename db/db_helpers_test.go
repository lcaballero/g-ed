package db

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func assertFile(t *testing.T, filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Logf("fail: file '%s' does not exist", filename)
		t.Fail()
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type testDbLifeCycle func(*DB)

func testDb(t *testing.T, name string) *DB {
	db, err := LoadFromFile(name)
	if err != nil {
		assert.Nil(t, err)
		t.Fail()
	}
	return db
}

func lifeCycleNamedDb(t *testing.T, name string, cycle testDbLifeCycle) {
	db := testDb(t, name)
	cycle(db)
	db.Close()
	assert.Nil(t, db.remove())
}

func lifeCycleDb(t *testing.T, cycle testDbLifeCycle) {
	name := RandStringRunes(20)
	lifeCycleNamedDb(t, name, cycle)
}
