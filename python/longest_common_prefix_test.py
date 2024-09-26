import unittest
from dataclasses import dataclass

from longest_common_prefix import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[str]
            expected: str

        return [
            Case(given=["flower", "flow", "flight"], expected="fl"),
            Case(given=["dog", "racecar", "car"], expected=""),
        ]

    def test_longest_common_prefix(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.longestCommonPrefix(case.given)
            assert actual == case.expected
