func gcdOfStrings(str1 string, str2 string) string {
    a := len(str1)
    b := len(str2)

    if str1 + str2 != str2 + str1 {
        return ""
    }  
    for b != 0 {
        nameoji := a % b
        a = b
        b = nameoji
    }
    return str1[:a]

}