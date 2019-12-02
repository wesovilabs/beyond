package helper

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CheckError(t *testing.T) {
	assert := assert.New(t)
	assert.NotPanics(func() { CheckError(nil) })
	assert.Panics(func() { CheckError(errors.New("error")) })
}
