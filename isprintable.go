package piscine

func IsPrintable(s string) bool {
	for _, is_printable := range s {
		if is_printable < 32 || is_printable > 126 {
			return false
		}
	}
	return true
}
