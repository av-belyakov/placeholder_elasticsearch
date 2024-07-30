package coremodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoreStorage(t *testing.T) {
	cs := newStorage()

	key := "testkey:case"
	cs.setRequest(key, "eventenrichment case")

	result, ok := cs.getRequest(key)
	assert.True(t, ok)
	assert.Contains(t, result, "eventenrichment case")

	key = "testkey:alert"
	cs.setRequest(key, "eventenrichment alert")

	result, ok = cs.getRequest(key)
	assert.True(t, ok)
	assert.Contains(t, result, "eventenrichment alert")
}
