package piscine

func IsLower(s string) bool {
	for _, is_lower := range s {
		if is_lower < 'a' || is_lower > 'z' {
			return false
		}
	}
	return true
}
