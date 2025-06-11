package piscine

func SortWordArr(a []string) {
	// Simple bubble sort for sorting the slice
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				// Swap the elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
