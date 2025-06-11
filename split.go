package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	sepLen := len(sep)
	sLen := len(s)

	for i := 0; i <= sLen-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	result = append(result, s[start:])
	return result
}
