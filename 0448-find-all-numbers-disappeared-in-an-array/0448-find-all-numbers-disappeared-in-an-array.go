func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    appeared := make([]bool, n+1)

    for _, num := range nums {
        appeared[num] = true
    }

    var result []int
    for i :=1;i<=n;i++{
        if !appeared[i] {
            result = append(result, i)
        }
    }
    return result
}