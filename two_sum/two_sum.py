class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        stored = {}
        for idx, num in enumerate(nums):
            diff = target - num
            if diff in stored:
                return [stored[diff], idx]
            stored[num] = idx
