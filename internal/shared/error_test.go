package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErr(t *testing.T) {
	err := NewErr("An error occurred")
	assert.NotNil(t, err)
	assert.Equal(t, "An error occurred", err.GetMessage())
	assert.Equal(t, 500, err.GetCode())
}

func TestErrorMethod(t *testing.T) {
	err := NewErr("Another error")
	assert.Equal(t, "Another error", err.Error())
}

func TestGetMessage(t *testing.T) {
	err := NewErr("Message retrieval")
	assert.Equal(t, "Message retrieval", err.GetMessage())
}

func TestGetCode(t *testing.T) {
	err := NewErr("Code retrieval")
	assert.Equal(t, 500, err.GetCode())
}

func TestSetCode(t *testing.T) {
	err := NewErr("Custom error")
	err.SetCode(404)
	assert.Equal(t, 404, err.GetCode())
}

func TestSetCodeReturnsSelf(t *testing.T) {
	err := NewErr("Self-returning error")
	result := err.SetCode(403)
	assert.Equal(t, err, result)
}
