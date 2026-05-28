func pivotInteger(n int) int {
    sum := n*(n+1)/2
    pivot := int(math.Sqrt(float64(sum)))

    if pivot * pivot == sum {
        return pivot
    }
    
    return -1
}