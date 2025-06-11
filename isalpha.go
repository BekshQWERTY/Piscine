package piscine

func IsAlpha(s string) bool {
	for _, is_alpha := range s {
		if !((is_alpha >= 'a' && is_alpha <= 'z') || (is_alpha >= 'A' && is_alpha <= 'Z') || (is_alpha >= '0' && is_alpha <= '9')) {
			return false
		}
	}
	return true
}
