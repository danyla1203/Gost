package test

import (
	. "../gost"
	"reflect"
	"testing"
)

func TestGetValuesFromUri(t *testing.T) {
	uri := []string{"user", "12"}
	pattern := "/user/:user_id"

	valuesFromUri := GetValuesFromUri(uri, pattern)
	expectedValues := map[string]int{"user_id": 12}
	t.Log("test uri =", uri, " pattern =", pattern)
	if !reflect.DeepEqual(valuesFromUri, expectedValues) {
		t.Error("Expected", expectedValues, "then got", valuesFromUri)
	}
}
