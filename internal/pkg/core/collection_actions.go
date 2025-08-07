package core

func FlattenErrors(toMap map[string]error) []error {
	result := make([]error, 0)
	for _, value := range toMap {
		result = append(result, value)
	}

	return result
}
