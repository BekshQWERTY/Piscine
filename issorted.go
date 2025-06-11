package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	ascendant, descendant := true, true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			descendant = false
		}
		if f(a[i], a[i+1]) > 0 {
			ascendant = false
		}
	}
	return ascendant || descendant
}
