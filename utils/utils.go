package utils

// Function to check if a value exists in a slice
func Contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}
