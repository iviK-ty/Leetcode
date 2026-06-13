func reverseWords(s string) string {
	rs := strings.Fields(s)
	left := 0
	right := len(rs) - 1

	for left < right {
		rs[left], rs[right] = rs[right], rs[left]
		left++
		right--
	}
	return strings.Join(rs, " ")
}