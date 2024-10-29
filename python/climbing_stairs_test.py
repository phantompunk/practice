import unittest
from dataclasses import dataclass

from climbing_stairs import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: int
            expected: int

        return [
            Case(given=2, expected=2),
            Case(given=3, expected=3),
            Case(given=4, expected=5),
        ]

    def test_climbing_stairs(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.climbStairs(case.given)  
            assert actual == case.expected  
