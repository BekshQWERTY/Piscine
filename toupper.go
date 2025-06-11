package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for i, to_upper := range result {
		if to_upper >= 'a' && to_upper <= 'z' {
			result[i] = to_upper - 32
		}
	}
	return string(result)
}
