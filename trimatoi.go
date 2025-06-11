package piscine

func TrimAtoi(s string) int {
	var result int
	sign := 1
	foundNum := false

	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			result = result*10 + int(letter-'0')
			foundNum = true
		} else if letter == '-' && !foundNum {
			sign = -1
		}
	}
	return result * sign
}
