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
	assert.Len(advices.List(), 6)
	assertAdvice(assert, advices.List()[0], `NewComplexAround("test",testAdvice.Attribute{})`, "github.com/wesovilabs/goa/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[1], "newEmptyReturning", "github.com/wesovilabs/goa/testdata", false, true)
	assertAdvice(assert, advices.List()[2], "newEmptyReturning", "github.com/wesovilabs/goa/testdata", false, true)
	assertAdvice(assert, advices.List()[3], "NewEmptyAround", "github.com/wesovilabs/goa/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[4], "NewTracingAdvice", "github.com/wesovilabs/goa/testdata/advice", true, false)
	assertAdvice(assert, advices.List()[5], "NewComplexBefore(&testAdvice.Attribute{})", "github.com/wesovilabs/goa/testdata/advice", true, false)

}

func assertAdvice(assert *assert.Assertions, advice *advice.Advice, name, pkg string, hasBefore, hasReturning bool) {
	assert.Equal(advice.Pkg(), pkg)
	assert.Equal(advice.Name(), name)
	assert.Equal(advice.HasBefore(), hasBefore)
	assert.Equal(advice.HasReturning(), hasReturning)
}
