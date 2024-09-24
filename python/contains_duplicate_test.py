import unittest
from dataclasses import dataclass

from contains_duplicate import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: bool

        return [
            Case(given=[1], expected=False),
            Case(given=[1,1], expected=True),
            Case(given=[1,2,3,4,5,9,5], expected=True),
            Case(given=[1,2,3,4,5,9,11], expected=False),
        ]

    def test_contains_duplicate(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.containsDuplicate(case.given)  
            assert actual == case.expected  
