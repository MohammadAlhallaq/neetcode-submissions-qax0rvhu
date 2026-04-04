func twoSum(nums []int, target int) []int {

    freq := make(map[int]int, len(nums))


	for k, v := range nums {
		com := target - v

		if idx, ok := freq[com]; ok {
			return []int{idx, k}
		}else{
			freq[v] = k
		}
	}

    return nil
	
}
