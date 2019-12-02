package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NormalizePointcut(t *testing.T) {
	assert := assert.New(t)
	expr := NormalizePointcut("pkg.Person?.fn(...)...")
	fmt.Println(expr.String())
	assert.True(expr.Match([]byte("pkg.Person.fn(string,*int)")))
	assert.True(expr.Match([]byte("pkg.fn(string,*int)")))
	assert.True(expr.Match([]byte("pkg.fn()")))
	assert.False(expr.Match([]byte("pkg.Object.fn(string,*int)")))
	assert.False(expr.Match([]byte("pkg.Person?.fn(string,*int)")))

	expr = NormalizePointcut("pkg.*?.fn(...)...")
	fmt.Println(expr.String())
	assert.True(expr.Match([]byte("pkg.Person.fn(string,*int)")))
	assert.True(expr.Match([]byte("pkg.fn(string,*int)")))
	assert.True(expr.Match([]byte("pkg.fn()")))
	assert.True(expr.Match([]byte("pkg.Object.fn(string,*int)")))
	assert.False(expr.Match([]byte("pkg.Person?.fn(string,*int)")))

}
