package packages4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	assert.Equal(t, "Hi,Gopher", getValue())
}
