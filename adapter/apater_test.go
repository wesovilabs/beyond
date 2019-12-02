package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasAnyReturning(t *testing.T) {
	assert := assert.New(t)
	assert.False(hasAnyReturning(nil))
}
