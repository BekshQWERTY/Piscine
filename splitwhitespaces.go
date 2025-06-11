package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var buffer []rune

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if len(buffer) > 0 {
				result = append(result, string(buffer))
				buffer = buffer[:0]
			}
		} else {
			buffer = append(buffer, r)
		}
	}

	if len(buffer) > 0 {
		result = append(result, string(buffer))
	}

	return result
}
