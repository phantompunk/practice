import unittest
from dataclasses import dataclass

from threesum import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: list[list[int]]

        return [
            Case(given=[0, 1, 1], expected=[]),
            Case(given=[0, 1, 1, -2], expected=[[-2, 1, 1]]),
            Case(given=[0, 0, 0], expected=[[0, 0, 0]]),
            Case(given=[-1, 0, 1, 2, -1, -4], expected=[[-1, -1, 2], [-1, 0, 1]]),
        ]

    def test_threesum(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.threeSum(case.given)
            assert actual == case.expected
