package lib

import (
	"reflect"
	"testing"
)

func TestGetValuesFromUri(t *testing.T) {
	uri := []string{"user", "12"}
	pattern := []string{"user", ":user_id"}

	valuesFromUri := getValuesFromUri(uri, pattern)
	expectedValues := map[string]int{"user_id": 12}
	t.Log("test uri =", uri, " pattern =", pattern)
	if !reflect.DeepEqual(valuesFromUri, expectedValues) {
		t.Error("Expected", expectedValues, "then got", valuesFromUri)
	}
}
