package testdata

import _package "github.com/wesovilabs/goa/function/testdata/package"

type person struct {
}

func noParams()                                                            {}
func singleParamString(param string)                                       {}
func singleParamPerson(param person)                                       {}
func singleParamPersonPointer(param *person)                               {}
func singleParamExternalPerson(param _package.Person)                      {}
func singleParamExternalPersonPointer(param *_package.Person)              {}
func singleParamInt32(param int32)                                         {}
func singleParamInt32Pointer(param *int32)                                 {}
func singleParamArrayOfString(param []string)                              {}
func singleParamArrayOfStringPointer(param []*string)                      {}
func singleParamInterface(interface{})                                     {}
func singleParamArrayOfInterface([]interface{})                            {}
func singleParamStruct(struct{})                                           {}
func singleParamArrayOfStruct([]struct{})                                  {}
func singleParamMapStringString(map[string]string)                         {}
func singleParamMap(map[string]string)                                     {}
func singleParamMapStringPerson(map[string]_package.Person)                {}
func singleParamMapStringPersonPointer(map[string]*_package.Person)        {}
func singleParamFuncEmpty(func())                                          {}
func singleParamFuncArg(func(*_package.Person))                            {}
func singleParamFuncArgStringPersonPointer(func(string, *_package.Person)) {}
func singleParamFuncArgStringInt(func(string, int))                        {}
func singleReturnFuncReturnStringInt() func(string, int)                   {}
