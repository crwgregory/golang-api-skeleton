package utils

import (
	"reflect"
	"testing"
)

func TestArrayContains(t *testing.T) {
	a := []string{
		"hello",
		"world",
	}
	o, e := ArrayContains(a, "hello", reflect.String)
	if e != nil {
		t.Fail()
	}
	if o != true {
		t.Fail()
	}
}
