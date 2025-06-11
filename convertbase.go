package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := 0
	baseFromLen := len(baseFrom)
	for _, r := range nbr {
		digit := 0
		for i, c := range baseFrom {
			if c == r {
				digit = i
				break
			}
		}
		decimal = decimal*baseFromLen + digit
	}

	if decimal == 0 {
		return string(baseTo[0])
	}
	var result []rune
	baseToLen := len(baseTo)
	for decimal > 0 {
		remaider := decimal % baseToLen
		result = append([]rune{rune(baseTo[remaider])}, result...)
		decimal /= baseToLen
	}
	return string(result)
}
