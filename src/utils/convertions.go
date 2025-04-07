package utils

import "strconv"

func AtoiOrDefault(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	if v, err := strconv.Atoi(value); err == nil {
		return v
	}
	return defaultValue
}
