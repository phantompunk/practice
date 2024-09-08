import unittest
from dataclasses import dataclass

from two_sum.two_sum import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            target: int
            expected: list[int]

        return [
            Case(given=[2, 7, 11, 15], target=9, expected=[0, 1]),
            Case(given=[3, 2, 4], target=6, expected=[1, 2]),
            Case(given=[3, 3], target=6, expected=[0, 1]),
            Case(given=[3, 6, 4, 2], target=8, expected=[1, 3]),
            Case(given=[3, 2, 4, 6], target=9, expected=[0, 3]),
        ]

    def test_two_sum(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.twoSum(case.given, case.target)
