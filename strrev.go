package piscine

func StrRev(s string) string {
	rev := ""
	for _, r := range s {
		rev = string(r) + rev
	}
	return rev
}
