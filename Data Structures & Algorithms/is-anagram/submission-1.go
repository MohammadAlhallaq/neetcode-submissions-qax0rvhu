func isAnagram(s string, t string) bool {

	if(len(s) != len(t)){
		return false;
	}

 	res := make(map[rune]int)

	for _, st := range s {
		res[st]++
	}


	for _, st := range t {
		res[st]--
		if(res[st] < 0){
			return false
		}
	}
	

	 return true
}
