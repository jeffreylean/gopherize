package variables1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariables(t *testing.T) {
	assert.Equal(t, Variables(), 1)
}
