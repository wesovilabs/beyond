package database

import (
        "fmt"
        "github.com/wesovilabs/goa/testdata/basic/model"
        "math/rand"
        "time"
)

var db = database{people: make(map[string]*model.Person)}

type database struct{ people map[string]*model.Person }

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func id() string {
        return fmt.Sprintf("%v", seededRand.Int())
}
