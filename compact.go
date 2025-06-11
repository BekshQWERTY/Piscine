package piscine

func Compact(ptr *[]string) int {
	temp := []string{}
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			temp = append(temp, (*ptr)[i])
		}
	}
	*ptr = []string{}
	for _, ch := range temp {
		*ptr = append(*ptr, ch)
	}
	return len(*ptr)
}
