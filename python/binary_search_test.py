import unittest
from dataclasses import dataclass

from binary_search import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            target: int
            expected: int

        return [
            Case(given=[3], target=3, expected=0),
            Case(given=[3, 4, 5, 6, 7], target=6, expected=3),
            Case(given=[-1, 0, 3, 5, 9, 12], target=4, expected=-1),
            Case(given=[-1, 0, 3, 5, 9, 12], target=12, expected=5),
        ]

    def test_longest_palindrome(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.search(case.given, case.target)
