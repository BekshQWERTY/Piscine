package piscine

func ReverseMenuIndex(menu []string) []string {
	// Create a new slice with the same length as the input slice
	reversed := make([]string, len(menu))

	// Iterate over the input slice in reverse order and copy elements to the new slice
	for i := range menu {
		reversed[i] = menu[len(menu)-1-i]
	}

	return reversed
}
