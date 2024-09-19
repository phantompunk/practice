class Solution(object):
    def longestPalindrome(self, s: str) -> int:
        """
        :type s: str
        :rtype: int
        """
        agg = 0
        hasMiddle = False
        for char in list(set(s)):
            if not hasMiddle and s.count(char) % 2 == 1:
                agg += 1
                hasMiddle = True
            if s.count(char) >= 2:
                agg = agg + (s.count(char)//2) * 2
        return agg
