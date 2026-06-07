func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := candies[0]
	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}
	result := make([]bool, len(candies))

	for i, candy2 := range candies {
		if candy2+extraCandies >= maxCandy {
			result[i] = true
		} else {
			result[i] = false
		}

	}
	fmt.Println(result)
	return result
}