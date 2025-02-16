package utils

import "strings"

func StringToSnakeCase(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 5) // Approximate additional space for underscores

	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}
