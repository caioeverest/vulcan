package console

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel_set(t *testing.T) {
	target := LevelDebug
	target.set(LevelWarn)
	assert.Equal(t, target, LevelWarn)
}

func TestLevel_permits(t *testing.T) {
	target := LevelWarn
	result := target.permits(LevelDebug)
	assert.False(t, result)

	result = target.permits(LevelFatal)
	assert.True(t, result)

	result = target.permits(LevelWarn)
	assert.True(t, result)
}
