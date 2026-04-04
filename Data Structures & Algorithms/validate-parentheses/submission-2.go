func isValid(s string) bool {
    
    stack := []rune{}

    for _, char := range s {
        if char == '{' || char == '[' || char == '(' {
            stack = append(stack, char)
        }else{
            if len(stack) == 0 {
    return false
}
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1] 
           if 
            (pop == '[' && char != ']') ||
            (pop == '{' && char != '}') ||
            (pop == '(' && char != ')') {
            return false
            }
        }
    }


    return len(stack) == 0
}
