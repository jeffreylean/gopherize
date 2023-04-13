package ifelse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigger(t *testing.T) {
	a, b := 33, 43

	assert.Equal(t, 43, bigger(a, b))
}
