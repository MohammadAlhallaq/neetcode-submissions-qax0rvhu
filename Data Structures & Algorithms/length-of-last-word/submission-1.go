func lengthOfLastWord(s string) int {
    count := 0
    i := len(s) - 1
	fmt.Print(i)

    for i >= 0 && s[i] == ' ' {
        i--
    }


    for i >= 0 && s[i] != ' ' {
        count++
        i--
    }

    return count
}