func canPlaceFlowers(flowerbed []int, n int) bool {
	planted := 0

	start := 0
	end := len(flowerbed) - 1

	for i := 0; i < len(flowerbed); i++ {

		if flowerbed[i] == 1 {
			continue
		}

		leftEmpty := (i == start || flowerbed[i-1] == 0)
		rightEmpty := (i == end || flowerbed[i+1] == 0)

		if leftEmpty && rightEmpty {
			flowerbed[i] = 1
			planted++
		}

		if planted == n {
			return true
		}
	}
	return planted >= n
}