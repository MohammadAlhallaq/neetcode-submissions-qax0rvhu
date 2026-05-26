class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
        const sMap = new Map<string, number>()

        if (s.length != t.length) {
            return false
        }

        for (let i = 0; i < s.length; i++) {
            const element = s[i];
            sMap.set(element, (sMap.get(element) || 0) + 1)
        }

        for (let j = 0; j < t.length; j++) {
            const element = t[j];
            if (sMap.get(element) || 0 > 0) {
                sMap.set(element, (sMap.get(element) || 0) - 1)
            } else {
                return false
            }
        }
        return true
    }
}







