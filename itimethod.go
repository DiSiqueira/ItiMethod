package itimethod

import (
	"net/http"
	"strings"
)

// MethodMatcher Store the template to match with the request
type MethodMatcher struct {
	methodList []string
}

// New is the constructor to ItiWildcard
func New(methodList ...string) *MethodMatcher {
	for i, v := range methodList {
		methodList[i] = strings.ToUpper(v)
	}
	return &MethodMatcher{methodList: methodList}
}

// Match if the request are using one of the acceptable methods.
func (m *MethodMatcher) Match(req *http.Request) bool {
	if len(m.methodList) <= 0 {
		return true
	}

	for _, v := range m.methodList {
		if req.Method == v {
			return true
		}
	}
	return false
}
