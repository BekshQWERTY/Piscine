package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i, to_lower := range result {
		if to_lower >= 'A' && to_lower <= 'Z' {
			result[i] = to_lower + 32
		}
	}
	return string(result)
}
