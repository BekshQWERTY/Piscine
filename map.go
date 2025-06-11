package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool

	for _, element := range a {
		result = append(result, f(element))
	}
	return result
}
