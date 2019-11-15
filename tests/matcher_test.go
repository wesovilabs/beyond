package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/goa/advice"
	"github.com/wesovilabs/goa/joinpoint"
	"github.com/wesovilabs/goa/match"
	"testing"
)

func Test_Matcher(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	advices := advice.GetAdvices(pkg, packages)
	joinPoints := joinpoint.GetJoinPoints(pkg, packages)
	matches := match.GetMatches(joinPoints, advices)
	assert.NotNil(matches)
	assert.Equal(10, len(matches))

}
