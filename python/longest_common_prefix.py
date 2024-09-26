from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        min_length = len(strs[0])
        for word in strs:
            min_length = min(len(word), min_length)

        index = 0
        prefix = ""
        match = True
        while index < min_length and match:
            char = ""
            for word in strs:
                if not char:
                    char = word[index : index + 1]
                if char != word[index : index + 1]:
                    match = False
                    break
            index += 1
            if match:
                prefix += char


        return prefix
