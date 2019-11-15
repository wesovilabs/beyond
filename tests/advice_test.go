package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/advice"
	"testing"
)

func Test_Advice(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	advices := advice.GetAdvices(pkg, packages)
	assert.NotNil(advices)
	assert.NotNil(advices.List())
	assert.Len(advices.List(), 4)
	assertAdvice(assert, advices.List()[0], `NewComplexAround`, "github.com/wesovilabs/goa/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[1], "NewEmptyAround", "github.com/wesovilabs/goa/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[2], "NewTracingAdvice", "github.com/wesovilabs/goa/api/advice", true, false)
	assertAdvice(assert, advices.List()[3], "NewComplexBefore", "github.com/wesovilabs/goa/testdata/advice", true, false)

}

func assertAdvice(assert *assert.Assertions, advice *advice.Advice, name, pkg string, hasBefore, hasReturning bool) {
	assert.Equal(advice.Pkg(), pkg)
	assert.Equal(advice.Name(), name)
	assert.Equal(advice.HasBefore(), hasBefore)
	assert.Equal(advice.HasReturning(), hasReturning)
}
