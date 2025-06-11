package piscine

func IsNumeric(s string) bool {
	for _, is_numeric := range s {
		if is_numeric < '0' || is_numeric > '9' {
			return false
		}
	}
	return true
}
