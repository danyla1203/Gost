package lib

func IsMiddleware(uri []string, pattern []string) bool {
	if len(uri) < len(pattern) {
		return false
	}
	if len(pattern) == 1 && len(pattern[0]) == 0 {
		return true
	}
	for index, _ := range pattern {
		if pattern[index] != uri[index] {
			return false
		}
	}
	return true
}
