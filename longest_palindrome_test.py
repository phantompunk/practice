import unittest
from dataclasses import dataclass

from longest_palindrome import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: str
            expected: int

        return [
            Case(given="a", expected=1),
            Case(given="abc", expected=1),
            Case(given="ccc", expected=3),
            Case(given="abccccdd", expected=7),
        ]

    def test_longest_palindrome(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.longestPalindrome(case.given)
