type Solution struct{}

func (s *Solution) Encode(strs []string) string {

    res := ""

    for _, str := range strs {
        res += strconv.Itoa(len(str)) + "#" + str

    }
    fmt.Println(res)  // adds a newline automatically

    return res

}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}
    i := 0
    for i < len(encoded) {
        // parse number before #
        j := i
        for j < len(encoded) && unicode.IsDigit(rune(encoded[j])) {
            j++
        }


        // parse the number
        numberStr := encoded[i:j]
        number, _ := strconv.Atoi(numberStr)
        start := j + 1
        end := start + number
        res = append(res, encoded[start:end])
        i = end
    }
    return res


}
