package utils

import (
	"errors"
	"reflect"
	"testing"
)

func isFunction(f interface{}) bool {
	actual := reflect.TypeOf(f)
	return actual.Kind() == reflect.Func && actual.NumIn() == 0 && actual.NumOut() > 0
}
func isChan(a interface{}) bool {
	if isNil(a) {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Chan
}
func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}
	return false
}
func getMapKeys(t *testing.T, m interface{}) []string {
	if reflect.ValueOf(m).Kind() != reflect.Map {
		t.Fatal(errors.New("GetMapKeys should receive a map"))
	}
	if reflect.TypeOf(m).Key() != reflect.TypeOf("bla") {
		t.Fatal(errors.New("GetMapKeys should receive a map with string keys"))
	}
	t.Helper()
	res := make([]string, 0)
	for _, k := range reflect.ValueOf(m).MapKeys() {
		res = append(res, k.String())
	}
	return res
}
