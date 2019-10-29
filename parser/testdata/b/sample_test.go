package b_test

import (
	"encoding/json"
	"github.com/wesovilabs/goa/parser/testdata/testutil"
)

func test(encoder json.Encoder) testutil.RandomIDS{
	return testutil.RandomIDS{}
}

func test2(encoder json.Encoder) *testutil.RandomIDS{
	return nil
}
