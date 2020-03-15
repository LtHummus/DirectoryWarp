package warps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSimple(t *testing.T) {
	empty := emptyWarps()

	assert.Len(t, empty.Warps, 0)
}

func TestHappyPath(t *testing.T) {
	db := emptyWarps()

	db.SetEntry("test", "/tmp")
	assert.Len(t, db.Warps, 1)

	entry, exists := db.GetEntry("test")
	assert.True(t, exists)

	assert.Equal(t, *entry, "/tmp")
}

func TestGetNonExist(t *testing.T) {
	db := emptyWarps()

	_, exists := db.GetEntry("bad")
	assert.False(t, exists)
}

func TestDelete(t *testing.T) {
	db := emptyWarps()
	db.SetEntry("test", "/tmp")

	entry, exists := db.GetEntry("test")
	assert.True(t, exists)
	assert.Equal(t, *entry, "/tmp")

	deleted := db.DeleteEntry("test")
	assert.True(t, deleted)

	assert.Len(t, db.Warps, 0)

	_, exists = db.GetEntry("test")
	assert.False(t, exists)
}

func TestDeleteNotExist(t *testing.T) {
	db := emptyWarps()

	deleted := db.DeleteEntry("bad")
	assert.False(t, deleted)
}