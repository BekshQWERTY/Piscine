package piscine

func CountIf(f func(string) bool, tap []string) int {
	count := 0
	for _, element := range tap {
		if f(element) {
			count++
		}
	}
	return count
}
