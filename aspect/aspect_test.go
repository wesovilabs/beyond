package aspect

import (
	"fmt"
	"go/parser"
	"go/token"
	"testing"
	"time"
)

func TestSearcher_Search(t *testing.T) {
	path := "testdata"
	//path:="/Users/ivan/Workspace/BBVA/ECSKERNEL/main-controller/internal/api"
	packages, err := parser.ParseDir(&token.FileSet{}, path, nil, parser.ParseComments)
	if err != nil {
		t.FailNow()
	}

	start2 := time.Now().Nanosecond()
	aspects := GetAspects(packages)
	fmt.Println(len(aspects.aroundList))
	end2 := time.Now().Nanosecond()
	fmt.Printf("%v ns\n", end2-start2)
}
