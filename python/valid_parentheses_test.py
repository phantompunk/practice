import unittest
from dataclasses import dataclass

from valid_parentheses import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: str
            expected: bool

        return [
            Case(given="()", expected=True),
            Case(given="(]", expected=False),
            Case(given="([])", expected=True),
            Case(given="()[]{}", expected=True),
            Case(given="(({[]}))", expected=True),
        ]

    def test_is_valid(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.isValid(case.given)
