import unittest
from dataclasses import dataclass

from valid_palindrome import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: str
            expected: bool

        return [
            Case(given=" ", expected=True),
            Case(given="abba", expected=True),
            Case(given="Ab,:c cba", expected=True),
            Case(given="101101", expected=True),
            Case(given="A man, a plan, a canal: Panama", expected=True),
            Case(given="race a car", expected=False),
        ]

    def test_is_palindrome(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.isPalindrome(case.given)
