class Solution(object):
    def canConstruct(self, ransomNote: str, magazine: str):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        if len(ransomNote) > len(magazine):
            return False

        for i in range(len(ransomNote)):
            if ransomNote[i] in magazine:
                magazine = magazine.replace(ransomNote[i], "", 1)
            else:
                return False
        return True

        if len(ransomNote) > len(magazine):
            return False

        freqMag = {}
        for c in magazine:
            freqMag[c] = 1 + freqMag.get(c, 0)

        for c in ransomNote:
            if c not in freqMag:
                return False
            else:
                freqMag[c] -= 1
                if freqMag[c] == 0:
                    del freqMag[c]
        return True
