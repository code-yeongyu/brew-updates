package main

import (
	"testing"

	"brew-updates/utils"

	"github.com/stretchr/testify/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
}

func TestExecuteCommand(t *testing.T) {
	result := utils.ExecuteCommand("echo hello")
	assert.Equal(t, "hello\n", result)
}
