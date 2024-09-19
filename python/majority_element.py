class Solution(object):
    def majorityElement(self, nums: list[int]) -> int:
        """
        :type nums: List[int]
        :rtype: int
        """
        result = majority = 0
        freq = {}
        for num in nums:
            freq[num] = 1 + freq.get(num, 0)
            if freq[num] >= majority:
                majority = freq[num]
                result = num
        return result
