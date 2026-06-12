func reverseVowels(s string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	bs := []byte(s)
	left := 0
	right := len(s) - 1

	for left < right {
		if !vowels[bs[left]] {
			left++
			fmt.Println("left++")
		} else if !vowels[bs[right]] {
			right--
			fmt.Println("right++")
		} else {
			bs[left], bs[right] = bs[right], bs[left]
			fmt.Println("changed")
			left++
			right--
		}
	}
	return string(bs)
}