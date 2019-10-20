package api

import "github.com/wesovilabs/goa/context"

type Settings struct {
	data map[string]interface{}
}

func (s *Settings) With(key string, value interface{}) *Settings {
	s.data[key] = value
	return s
}

func (s *Settings) GetString(key string) string {
	if v, ok := s.data[key]; ok {
		return v.(string)
	}
	return ""
}

type Aspect interface {
	SetUp(settings *Settings) Aspect
}

type Around interface {
	Aspect
	Around(string, string, *context.Input)
}
