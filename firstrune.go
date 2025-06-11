package piscine

func FirstRune(s string) rune {
	result := rune(' ')
	for _, value := range s {
		result = rune(value)
		break
	}
	return result
}
