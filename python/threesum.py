from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        print(nums)
        result = []

        for i in range(len(nums)):
            j = i + 1
            k = len(nums) - 1
            while j < k:
                total = nums[i] + nums[j] + nums[k]
                print(f"Nums:{i, nums[i]}, {nums[j]}, {nums[k]}, Sum:{total}")
                if total == 0:
                    if [nums[i], nums[j], nums[k]] not in result:
                        result.append([nums[i], nums[j], nums[k]])
                if total <= 0:
                    j += 1
                if total >= 0:
                    k -= 1

        return result
