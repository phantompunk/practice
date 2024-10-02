import unittest
from dataclasses import dataclass

from insert_interval import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[list[int]]
            new: list[int]
            expected: list[list[int]]

        return [
            Case(given=[[1, 2], [3, 4]], new=[5, 6], expected=[[1, 2], [3, 4], [5, 6]]),
            Case(given=[[1, 2], [8, 9]], new=[5, 6], expected=[[1, 2], [5, 6], [8, 9]]),
            Case(given=[[1, 2], [3, 4]], new=[2, 6], expected=[[1, 6]]),
            Case(given=[[1, 3], [6, 9]], new=[2, 5], expected=[[1, 5], [6, 9]]),
            Case(given=[[1, 2], [4, 5]], new=[2, 4], expected=[[1, 5]]),
        ]

    def test_insert_interval(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.insert(case.given, case.new)
            assert actual == case.expected
