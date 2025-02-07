from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        group = dict()

        for s in strs:
            key = "".join(sorted(s))
            if key in group:
                group[key].append(s)
            else:
                group[key] = [s]

        result = []
        for key in group:
            result.append(group[key])

        return result

