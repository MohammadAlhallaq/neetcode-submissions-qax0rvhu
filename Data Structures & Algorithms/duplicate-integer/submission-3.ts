class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
  hasDuplicate(nums: number[]): boolean {
        const scores = new Map<number, boolean>()
        for (let index = 0; index < nums.length; index++) {
            if (scores.has(nums[index])) {
                return true
            } else {
                scores.set(nums[index], true)
            }
        }
        return false
    }
}
