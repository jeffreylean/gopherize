package for2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	assert.Equal(t, 10, loop())
}
