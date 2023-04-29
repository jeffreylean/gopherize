package switch1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	assert.Equal(t, true, isEven(2))
	assert.Equal(t, false, isEven(3))
}
