package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
}

func TestExecuteCommand(t *testing.T) {
	result := ExecuteCommand("echo hello")
	assert.Equal(t, "hello\n", result)
}
