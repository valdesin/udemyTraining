package stringutil

func reverseTwo(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i++ {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
