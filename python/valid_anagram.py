class Solution(object):
    def isAnagram(self, s: str, t: str):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if len(s) != len(t):
            return False

        mappingS = {}
        mappingT = {}
        for i in range(len(s)):
            mappingS[s[i]] = 1 + mappingS.get(s[i], 0)
            mappingT[t[i]] = 1 + mappingT.get(t[i], 0)
        return mappingS == mappingT
