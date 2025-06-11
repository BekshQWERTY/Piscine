package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)

	if n == 0 {
		return podium
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
