package test

import (
	. "../gost"
	"testing"
)

func test(testModule *testing.T, uri []string, pattern []string) func(toBe bool) {
	testModule.Log("test uri =", uri, " pattern =", pattern)
	result := CheckPath(uri, pattern)
	return func(toBe bool) {
		if result != toBe {
			testModule.Error("must be", toBe, "then got", result)
		}
	}
}

func TestWithHardcodedPattern(t *testing.T) {
	uri := []string{""}
	pattern := []string{""}
	test(t, uri, pattern)(true)

	uri = []string{"docs", "12"}
	pattern = []string{"docs"}
	test(t, uri, pattern)(false)
}

func TestWithRegexpAndVars(t *testing.T) {
	uri := []string{"user", "1423"}
	pattern := []string{"user", ":user_id"}
	test(t, uri, pattern)(true)

	uri = []string{"sdklfj234", "22"}
	pattern = []string{"*", ":id"}
	test(t, uri, pattern)(true)

	uri = []string{"randWOrd", "22", "randWord"}
	pattern = []string{"*", ":number", "*"}
	test(t, uri, pattern)(true)
}

func TestWithComplexRegexp(t *testing.T) {
	uri := []string{"user", "user-1234"}
	pattern := []string{"user", "\\w+-\\d+"}
	test(t, uri, pattern)(true)
}
