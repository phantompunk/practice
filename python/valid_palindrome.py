class Solution(object):
    def isPalindrome(self, s: str):
        """
        :type s: str
        :rtype: bool
        """
        # mod = s.lower().replace(" ", "").replace(",", "").replace(":", "")
        # return "".join(mod[::-1]) == mod

        i = 0
        j = len(s) - 1
        s = s.lower()
        while i < j:
            if s[i] < "a" or s[i] > "z":
                i += 1
            elif s[j] < "a" or s[j] > "z":
                j -= 1
            elif s[i] == s[j]:
                i += 1
                j -= 1
            else:
                return False
        return True
