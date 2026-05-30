func mergeAlternately(word1 string, word2 string) string {
    byteArr := []byte{}
    maxLen := len(word1)
    if maxLen < len(word2) {
        maxLen = len(word2)
    }

    for i:=0;i<maxLen;i++{
        if i < len(word1) {
            byteArr = append(byteArr, word1[i])

        }
        if  i < len(word2) {
            byteArr = append(byteArr, word2[i])
        }
    }

    return string(byteArr)
}