package test

import (
	. "../gost"
	"testing"
)

func testMiddleware(testModule *testing.T, uri []string, pattern []string) func(toBe bool) {
	testModule.Log("test uri =", uri, " pattern =", pattern)
	result := IsMiddleware(uri, pattern)
	return func(toBe bool) {
		if result != toBe {
			testModule.Error("must be", toBe, "then got", result)
		}
	}
}

func TestSimpMiddlewares(t *testing.T) {
	uri := []string{"article", "1234"}
	pattern := []string{"article"}
	testMiddleware(t, uri, pattern)(true)

	uri = []string{"user", "1234"}
	pattern = []string{""}
	testMiddleware(t, uri, pattern)(true)

	uri = []string{"article", "admin", "1234"}
	pattern = []string{"article", "admin"}
	testMiddleware(t, uri, pattern)(true)

	uri = []string{"article", "admin", "1234"}
	pattern = []string{"admin"}
	testMiddleware(t, uri, pattern)(false)
}
