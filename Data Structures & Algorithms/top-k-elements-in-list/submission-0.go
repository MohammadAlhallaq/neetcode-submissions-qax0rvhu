func topKFrequent(nums []int, k int) []int {


	freq := make(map[int]int, len(nums))

	for _, v := range nums {
		freq[v]++
	}

	n := len(nums)
    buckets := make([][]int, n+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}


    result := []int{}
	for f := n; f >= 0 && k > len(result); f-- {
		if len(buckets[f]) > 0 {
            result = append(result, buckets[f]...)
        }
	}

	return result[:k]

}
