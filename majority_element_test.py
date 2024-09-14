import unittest
from dataclasses import dataclass

from majority_element import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: int

        return [
            Case(given=[1], expected=1),
            Case(given=[3, 2, 3], expected=3),
            Case(given=[2, 2, 1, 1, 1, 2, 2], expected=2),
        ]

    def test_majorityElement(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.majorityElement(case.given)
