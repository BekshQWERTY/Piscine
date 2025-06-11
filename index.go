package piscine

func Index(s string, toFind string) int {
	n := len(toFind)
	switch {
	case n == 0:
		return 0
	case n > len(s):
		return -1
	}
	for i := 0; i <= len(s)-n; i++ {
		if s[i:i+n] == toFind {
			return i
		}
	}
	return -1
}
