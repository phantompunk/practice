class Solution(object):
    def isPalindrome(self, s: str):
        """
        :type s: str
        :rtype: bool
        """
        mod = s.lower().replace(" ", "").replace(",", "").replace(":", "")
        return ''.join(reversed(mod)) == mod
