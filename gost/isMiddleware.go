package lib

func IsMiddleware(uri []string, pattern []string) bool {
	if len(uri) < len(pattern) {
		return false
	}

	for index, _ := range pattern {
		if pattern[index] != uri[index] {
			return false
		}
	}
	return true
}
