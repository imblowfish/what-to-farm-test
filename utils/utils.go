package utils

func Contains(requiredValues []string, v string) bool {
	for _, reqValue := range requiredValues {
		if v == reqValue {
			return true
		}
	}
	return false
}
