package cmd

func contains(requiredValues []string, v string) bool {
	for _, reqValue := range requiredValues {
		if v == reqValue {
			return true
		}
	}
	return false
}
