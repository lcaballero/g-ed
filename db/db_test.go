package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cannot_Save_To_Empty_Key(t *testing.T) {
	lifeCycleDb(t, func(db *DB) {
		assert.NotNil(t, db.data)
	})
}

func Test_When_No_Exists(t *testing.T) {
	name := "when-not-exists.db"
	lifeCycleNamedDb(t, name, func(db *DB) {
		assertFile(t, name)
	})
}

func Test_Write_Key(t *testing.T) {
	lifeCycleDb(t, func(db *DB) {
		writeErr := db.Set("val1", "key1")
		assert.Nil(t, writeErr)
		assert.True(t, db.Has("val1"))
	})
}

func Test_Read_Values(t *testing.T) {
	name := "on-re-open.db"

	db := testDb(t, name)
	db.Set("val2", "key2")
	assert.True(t, db.Has("val2"), "should have found for saved key")
	db.Close()

	db = testDb(t, name)
	assert.True(t, db.Has("val2"), "should have found for saved key")
	db.Close()

	assert.Nil(t, db.remove())
}
