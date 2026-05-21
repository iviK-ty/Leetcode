func findMaxConsecutiveOnes(nums []int) int {
    resultCount := 0
    currentCount := 0

    for _, num := range nums {
        if num == 1{
            currentCount++

           if currentCount > resultCount {
                resultCount = currentCount
           } 
           }else{
            currentCount = 0
        }
    }    
    return resultCount
}