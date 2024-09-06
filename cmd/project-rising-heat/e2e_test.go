package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		version = "test-version"
		res, err := executeSubCommand("version")
		assert.NoError(t, err)
		assert.Equal(t, res, version+"\n")
	})
}
